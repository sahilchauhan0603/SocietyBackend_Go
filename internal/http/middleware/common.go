package middleware

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type contextKey string

const requestIDContextKey contextKey = "request_id"

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
	size       int
}

type captureResponseWriter struct {
	headers    http.Header
	statusCode int
	body       bytes.Buffer
}

func (crw *captureResponseWriter) Header() http.Header {
	return crw.headers
}

func (crw *captureResponseWriter) WriteHeader(statusCode int) {
	crw.statusCode = statusCode
}

func (crw *captureResponseWriter) Write(body []byte) (int, error) {
	if crw.statusCode == 0 {
		crw.statusCode = http.StatusOK
	}
	return crw.body.Write(body)
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.statusCode = code
	sr.ResponseWriter.WriteHeader(code)
}

func (sr *statusRecorder) Write(body []byte) (int, error) {
	if sr.statusCode == 0 {
		sr.statusCode = http.StatusOK
	}
	n, err := sr.ResponseWriter.Write(body)
	sr.size += n
	return n, err
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		w.Header().Set("X-Request-ID", requestID)
		ctx := context.WithValue(r.Context(), requestIDContextKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDContextKey).(string)
	return requestID
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w}
		next.ServeHTTP(recorder, r)

		requestID := GetRequestID(r.Context())
		if recorder.statusCode == 0 {
			recorder.statusCode = http.StatusOK
		}

		entry := map[string]any{
			"timestamp":   time.Now().UTC().Format(time.RFC3339Nano),
			"level":       "info",
			"request_id":  requestID,
			"method":      r.Method,
			"path":        r.URL.Path,
			"status":      recorder.statusCode,
			"duration_ms": time.Since(start).Milliseconds(),
			"bytes":       recorder.size,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
		}

		payload, err := json.Marshal(entry)
		if err != nil {
			log.Printf("{\"level\":\"error\",\"message\":\"failed to marshal log entry\",\"error\":%q}", err.Error())
			return
		}

		log.Print(string(payload))
	})
}

func TraceResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capture := &captureResponseWriter{headers: make(http.Header)}
		next.ServeHTTP(capture, r)

		statusCode := capture.statusCode
		if statusCode == 0 {
			statusCode = http.StatusOK
		}

		requestID := GetRequestID(r.Context())
		body := capture.body.Bytes()

		for key, values := range capture.headers {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		if len(body) > 0 {
			if isJSONResponse(capture.headers.Get("Content-Type"), body) {
				updated, ok := appendTraceID(body, requestID)
				if ok {
					body = updated
				}
			}

			w.Header().Set("Content-Length", strconv.Itoa(len(body)))
		}

		w.WriteHeader(statusCode)
		if len(body) > 0 {
			_, _ = w.Write(body)
		}
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Printf("panic recovered: %v", recovered)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(map[string]string{
					"error": "internal server error",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	allowed := make(map[string]struct{}, len(allowedOrigins))
	allowAll := false
	for _, origin := range allowedOrigins {
		clean := strings.TrimSpace(origin)
		if clean == "*" {
			allowAll = true
		}
		if clean != "" {
			allowed[clean] = struct{}{}
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" {
				_, ok := allowed[origin]
				if allowAll || ok {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Request-ID")
					w.Header().Set("Vary", "Origin")
				}
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func generateRequestID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return time.Now().Format("20060102150405.000000")
	}
	return hex.EncodeToString(bytes)
}

func isJSONResponse(contentType string, body []byte) bool {
	if strings.Contains(strings.ToLower(contentType), "application/json") {
		return true
	}

	trimmed := bytes.TrimSpace(body)
	if len(trimmed) == 0 {
		return false
	}

	return trimmed[0] == '{' || trimmed[0] == '['
}

func appendTraceID(body []byte, requestID string) ([]byte, bool) {
	if requestID == "" {
		return body, false
	}

	var payload any
	if err := json.Unmarshal(body, &payload); err != nil {
		return body, false
	}

	switch typed := payload.(type) {
	case map[string]any:
		if _, exists := typed["trace_id"]; !exists {
			typed["trace_id"] = requestID
		}
		updated, err := json.Marshal(typed)
		if err != nil {
			return body, false
		}
		return updated, true
	case []any:
		wrapped := map[string]any{
			"trace_id": requestID,
			"data":     typed,
		}
		updated, err := json.Marshal(wrapped)
		if err != nil {
			return body, false
		}
		return updated, true
	default:
		return body, false
	}
}
