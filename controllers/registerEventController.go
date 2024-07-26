package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sahilchauhan0603/society/helper"
)

type RegisterEvent struct {
	FullName     string
	Email        string
	Batch        string
	Branch       string
	EnrollmentNo int64
}

func RegisterEventHandler(w http.ResponseWriter, r *http.Request) {

	emailUser := os.Getenv("EMAIL_USER")

	var formdata RegisterEvent
	if err := json.NewDecoder(r.Body).Decode(&formdata); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	emailBody := fmt.Sprintf("<p> FullName : %s </p> </br> <p>Email: %s</p> <p>Batch: %s</p> <p>Branch: %s</p> <p>EnrollmentNo : %d</p>", formdata.FullName, formdata.Email, formdata.Batch, formdata.Branch, formdata.EnrollmentNo)
	err := helper.SendEmail(emailUser, "EVENT REGISTER FORM", emailBody)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "mail sent successfully",
	})
}
