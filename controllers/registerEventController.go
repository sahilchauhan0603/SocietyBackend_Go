package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sahilchauhan0603/society/helper"
)

type RegisterEvent struct {
	Name       string
	Email      string
	Batch      string
	Branch     string
	Enrollment int64
	Society    string
}

var societyEmail = map[string]string{
	"Namespace":   "jyoti43cseb22@bpitindia.edu.in",
	"Anveshan":    "parth83cseb22@bpitindia.edu.in",
	"Hash Define": "mohit84cseb22@bpitindia.edu.in",
	"WIBD":        "jyoti43cseb22@bpitindia.edu.in",
	"GDSC":        "sahil82cseb22@bpitindia.edu.in",
	"WIE":         "jyoti43cseb22@bpitindia.edu.in",
	"IEEE":        "sahil82cseb22@bpitindia.edu.in",
	"Electonauts": "harsh63cseb22@bpitindia.edu.in",
	"Dhrishti":    "jyoti43cseb22@bpitindia.edu.in",
	"Opti Click":  "tanmay59cseb22@bpitindia.edu.in",
	"Avaran":      "tanmay59cseb22@bpitindia.edu.in",
	"Octave":      "ritesh100cseb22@bpitindia.edu.in",
	"Panache":     "sahil82cseb22@bpitindia.edu.in",	
	"Mavericks":   "sahil82cseb22@bpitindia.edu.in",
	"Kalam":       "sahil82cseb22@bpitindia.edu.in",
	"Chromavita":  "sahil82cseb22@bpitindia.edu.in",
}

// Handler to process the society query form submission
func RegisterForEvent(w http.ResponseWriter, r *http.Request) {

	var formdata RegisterEvent
	if err := json.NewDecoder(r.Body).Decode(&formdata); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find the email address for the selected society
	societyEmail, exists := societyEmail[formdata.Society]
	if !exists {
		http.Error(w, "Invalid society selected", http.StatusBadRequest)
		return
	}

	//email body
	emailBody := fmt.Sprintf(`<!DOCTYPE html>
    <html>
    <head>
        <title>Event Registration Form Submission</title>
        <style>
            body {
                font-family: Arial, sans-serif;
            }
            .container {
                padding: 20px;
            }
            .header {
                background-color: #f8f9fa;
                padding: 10px 20px;
                border-bottom: 1px solid #dee2e6;
            }
            .content {
                margin: 20px 0;
            }
            .footer {
                background-color: #f8f9fa;
                padding: 10px 20px;
                border-top: 1px solid #dee2e6;
                text-align: center;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <div class="header">
                <h2>I want to register for the event that your society is conducting</h2>
            </div>
            <div class="content">
                <p><strong>Student Name:</strong> %s </p>
                <p><strong>Enrollment No:</strong> %d</p>
                <p><strong>Branch:</strong> %s</p>
                <p><strong>Batch Year:</strong> %s</p>
                <p><strong>Email:</strong> %s</p>
                <p><strong>Society To Join:</strong> %s</p>
            </div>
            <div class="footer">
                <p>This message was sent from the BPIT Society Management Event Form.</p>
            </div>
        </div>
    </body>
    </html>`,
		formdata.Name,
		formdata.Enrollment,
		formdata.Branch,
		formdata.Batch,
		formdata.Email,
		formdata.Society)

	// Send email to the respective society email
	err := helper.SendEmail(societyEmail, "New Form Submitted for Event Registration from BPIT Society Management Portal", emailBody)
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

