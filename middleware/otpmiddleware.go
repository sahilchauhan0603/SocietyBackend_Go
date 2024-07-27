package middleware

import (
	"encoding/json"
	"github.com/sahilchauhan0603/society/helper"
	"net/http"
)

func OTPVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.Header.Get("Email")
		otp := r.Header.Get("OTP")

		if email == "" || otp == "" {
			http.Error(w, "Email and OTP are required", http.StatusBadRequest)
			return
		}
		err := helper.VerifyOTP(email, otp)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
