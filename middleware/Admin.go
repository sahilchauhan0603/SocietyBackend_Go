package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sahilchauhan0603/society/controllers"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Printf("Error retrieving token cookie: %v", err)
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &controllers.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			log.Println("Invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Check token expiration
		if float64(time.Now().Unix()) > float64(claims.ExpiresAt) {
			log.Println("Token expired")
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		if claims.Role != "admin" {
			log.Println("Access denied: user is not an admin")
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
