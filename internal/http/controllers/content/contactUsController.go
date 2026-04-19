package content

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sahilchauhan0603/society/internal/helpers"
)

type Contactform struct {
	Name      string
	Email     string
	ContactNo string
	Batch     string
	Branch    string
	Society   string
	Subject   string
	Message   string
}

func ContactUSHandler(w http.ResponseWriter, r *http.Request) {

	emailUser := os.Getenv("EMAIL_USER")

	var data Contactform
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	emailBody := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>New Contact Us Form Submission</title>
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
            <h2>New Contact Us Form Submission</h2>
        </div>
        <div class="content">
            <p><strong>Name:</strong> %s</p>
            <p><strong>Email:</strong> %s</p>
            <p><strong>ContactNo:</strong> %s</p>
            <p><strong>Batch:</strong> %s</p>
            <p><strong>Branch:</strong> %s</p>
            <p><strong>Society:</strong> %s</p>
            <p><strong>Subject:</strong> %s</p>
            <p><strong>Message:</strong></p>
            <p>%s</p>
        </div>
        <div class="footer">
            <p>This message was sent from the BPIT Society Management Contact Us form.</p>
        </div>
    </div>
</body>
</html>`, data.Name, data.Email, data.ContactNo, data.Batch, data.Branch, data.Society, data.Subject, data.Message)
	err := helper.SendEmail(emailUser, "New Contact Us Form Submission from BPIT Society Management Portal", emailBody)
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
