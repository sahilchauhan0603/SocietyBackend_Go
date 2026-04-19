package routes

import (
	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/society/internal/http/middleware"
	adminctrl "github.com/sahilchauhan0603/society/internal/service/admin"
)

func registerAdminRoutes(root *mux.Router, api *mux.Router) {
	root.HandleFunc("/adminlogin", adminctrl.AdminLogin).Methods("POST")

	adminRouter := root.PathPrefix("/api/v1/admin").Subrouter()
	adminRouter.Use(middleware.AdminMiddleware)
	adminRouter.HandleFunc("/adminrole", adminctrl.AddNewAdminRole).Methods("POST")
	adminRouter.HandleFunc("/adminrole/{username}", adminctrl.RemoveAdminRole).Methods("DELETE")

	api.HandleFunc("/adminroles", adminctrl.FetchAllAdminRoles).Methods("GET")
	api.HandleFunc("/adminrole/{username}", adminctrl.FetchAdminRole).Methods("GET")
	api.HandleFunc("/adminrole/{username}", adminctrl.UpdateAdminRole).Methods("PUT")
}
