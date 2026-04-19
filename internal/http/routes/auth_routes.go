package routes

import (
	"github.com/gorilla/mux"
	authctrl "github.com/sahilchauhan0603/society/internal/service/auth"
)

func registerAuthRoutes(root *mux.Router, api *mux.Router) {
	root.HandleFunc("/microsoftLogin", authctrl.HandleMicrosoftLogin).Methods("GET")
	root.HandleFunc("/callback", authctrl.HandleMicrosoftCallback).Methods("GET")
	root.HandleFunc("/forgotPassword", authctrl.SendEmail).Methods("POST")
	root.HandleFunc("/resetPassword", authctrl.VerifyReset).Methods("POST")

	api.HandleFunc("/login", authctrl.Login).Methods("POST")
	api.HandleFunc("/signup", authctrl.Signup).Methods("POST")
}
