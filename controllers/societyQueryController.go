package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sahilchauhan0603/society/helper"
)

type SocietyQuery struct {
	StudentName             string
	Society                 string
	Batch                   string
	Branch                  string
	StudentEnrollmentNumber int64
	Query                   string
}

// Map of society names/IDs to their respective email addresses
var societyEmails = map[string]string{
	"Namespace":   "namespace@society.com",
	"Anveshan":    "anveshan@society.com",
	"Hash Define": "hashdefine@society.com",
	"WIBD":        "wibd@society.com",
	"GDSC":        "gdsc@society.com",
	// Add more societies as needed
}

// func SocietyQueryHandler(w http.ResponseWriter, r *http.Request) {

// 	emailUser := os.Getenv("EMAIL_USER")

// 	var formdata SocietyQuery
// 	if err := json.NewDecoder(r.Body).Decode(&formdata); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	emailBody := fmt.Sprintf("<p> StudentName : %s </p> <p>Society: %s</p> <p>Batch: %s</p> <p>Branch: %s</p> <p>EnrollmentNo : %d</p> <p>Query : %s</p>", formdata.StudentName, formdata.Society, formdata.Batch, formdata.Branch, formdata.StudentEnrollmentNumber, formdata.Query)
// 	err := helper.SendEmail(emailUser, "SOCIETY QUERY FORM", emailBody)
// 	if err != nil {
// 		log.Printf("Error sending email: %v\n", err)
// 		http.Error(w, "Failed to send email", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "mail sent successfully",
// 	})
// }

// Handler to process the society query form submission
func SocietyQueryHandler(w http.ResponseWriter, r *http.Request) {

	var formdata SocietyQuery
	if err := json.NewDecoder(r.Body).Decode(&formdata); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the email address for the selected society
	societyEmail, exists := societyEmails[formdata.Society]
	if !exists {
		http.Error(w, "Invalid society selected", http.StatusBadRequest)
		return
	}

	// Create the email body
	emailBody := fmt.Sprintf("<p>Student Name: %s</p><p>Society: %s</p><p>Batch: %s</p><p>Branch: %s</p><p>Enrollment No: %d</p><p>Query: %s</p>", formdata.StudentName, formdata.Society, formdata.Batch, formdata.Branch, formdata.StudentEnrollmentNumber, formdata.Query)

	// Send email to the respective society email
	err := helper.SendEmail(societyEmail, "SOCIETY QUERY FORM", emailBody)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Mail sent successfully to " + societyEmail,
	})
}
