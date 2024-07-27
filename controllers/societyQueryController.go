package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sahilchauhan0603/society/helper"
)

type SocietyQuery struct {
	StudentName  string
	Society      string
	Batch        string
	Branch       string
	EnrollmentNo int64
	Query        string
}

func SocietyQueryHandler(w http.ResponseWriter, r *http.Request) {

	emailUser := os.Getenv("EMAIL_USER")

	var formdata SocietyQuery
	if err := json.NewDecoder(r.Body).Decode(&formdata); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	emailBody := fmt.Sprintf("<p> StudentName : %s </p> <p>Society: %s</p> <p>Batch: %s</p> <p>Branch: %s</p> <p>EnrollmentNo : %d</p> <p>Query : %s</p>", formdata.StudentName, formdata.Society, formdata.Batch, formdata.Branch, formdata.EnrollmentNo, formdata.Query)
	err := helper.SendEmail(emailUser, "SOCIETY QUERY FORM", emailBody)
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
