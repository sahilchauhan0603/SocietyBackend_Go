package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sahilchauhan0603/society/helper"
)

type CreateSociety struct {
	SocietyName        string
	HeadName           string
	DateOfRegistration string
	SocietyImage       string
	Category           string
	MobileNo           string
	Email              string
	Website            string
	Describe           string
}

func CreateSocietyHandler(w http.ResponseWriter, r *http.Request) {

	emailUser := os.Getenv("EMAIL_USER")

	var data CreateSociety
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	emailBody := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Create Society Form Submission</title>
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
            <h2>New Society Registration Request Form</h2>
        </div>
        <div class="content">
            <p><strong>Society Name:</strong> %s</p>
			<h3>Society Head Details:</h3>
                <p><strong>Society Head Name:</strong> %s</p>
                <p><strong>Society Head Email:</strong> %s</p>
                <p><strong>Society Head ContactNo:</strong> %s</p>
            <p><strong>Society's Website:</strong> %s</p>
            <p><strong>Society's Category:</strong> %s</p>
            <p><strong>Society's Date of Registration:</strong> %s</p>
            <p><strong>Society's Description:</strong></p>
            <p>%s</p>
        </div>
        <div class="footer">
            <p>This message was sent from the BPIT Society Management Contact Us form.</p>
        </div>
    </div>
</body>
</html>`, data.SocietyName, data.HeadName, data.Email, data.MobileNo, data.Website, data.Category, data.DateOfRegistration, data.Describe)
	err := helper.SendEmail(emailUser, "New Society Registration Request Form Submission from BPIT Society Management Portal", emailBody)
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
