package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sahilchauhan0603/society/helper"
)

type Contactform struct {
	name    string
	email   string 
	message string 
}

func ContactUSHandler(w http.ResponseWriter, r *http.Request) {
	var formdata Contactform
	if err := json.NewDecoder(r.Body).Decode(&formdata); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	emailBody := fmt.Sprintf("<h1> Name : %s </h1> </br> <p>email: %s</p> <p>message : %s</p>", formdata.name, formdata.email, formdata.message)
	err := helper.SendEmail("mohit84cseb22@bpitindia.edu.in", "CONTACT FORM", emailBody)
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
