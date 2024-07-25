package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sahilchauhan0603/society/helper"
)

type Contactform struct {
	Name    string
	Email   string
	ContactNo string
    Batch string
    Branch string
    Society string
	Subject string
	Message string
}

func ContactUSHandler(w http.ResponseWriter, r *http.Request) {

	emailUser := os.Getenv("EMAIL_USER")

	var formdata Contactform
	if err := json.NewDecoder(r.Body).Decode(&formdata); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	emailBody := fmt.Sprintf("<h1> Name : %s </h1> </br> <p>Email: %s</p> <p>ContactNo: %s</p> <p>Batch: %s</p> <p>Branch: %s</p> <p>Society : %s</p> <p>Subject: %s</p> <p>Message: %s</p>", formdata.Name, formdata.Email, formdata.ContactNo, formdata.Batch, formdata.Branch, formdata.Society, formdata.Subject, formdata.Message)
	err := helper.SendEmail(emailUser, "CONTACT FORM", emailBody)
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
