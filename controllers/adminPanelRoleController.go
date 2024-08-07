package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/society/config"
	models "github.com/sahilchauhan0603/society/models"
	"golang.org/x/crypto/bcrypt"
)

func AddNewAdminRole(w http.ResponseWriter, r *http.Request) {
	// Log the request
	log.Println("Received request to add new admin role")
	var role models.AdminPanelRole
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Log the decoded role
	log.Printf("Decoded role: %+v", role)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(role.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to hash password: ", err)
	}
	role.Password = string(hashedPassword)
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AdminPanelRole{}) {
		log.Println("AdminPanelRole table does not exist, creating it")
		if err := database.DB.AutoMigrate(&models.AdminPanelRole{}); err != nil {
			log.Printf("Error creating AdminPanelRole table: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := database.DB.Create(&role).Error; err != nil {
		log.Printf("Error inserting new role into database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
	log.Println("Successfully added new admin role")
}

func FetchAllAdminRoles(w http.ResponseWriter, r *http.Request) {

	var roles []models.AdminPanelRole
	if err := database.DB.Find(&roles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

func FetchAdminRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	var role []models.AdminPanelRole
	if err := database.DB.First(&role, "username = ?", username).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func UpdateAdminRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	var role models.AdminPanelRole
	if err := database.DB.First(&role, "username = ?", username).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&role).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(role)
}

func RemoveAdminRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	if err := database.DB.Delete(&models.AdminPanelRole{}, "username = ?", username).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Role successfully deleted"})
}

// Admin Login

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var admin models.AdminPanelRole
	if err := database.DB.Where("username = ?", creds.Username).First(&admin).Error; err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 72)
	claims := &Claims{
		Username: creds.Username,
		Role:     admin.Role,
		ID:       admin.SocietyID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   tokenString,
	})
}
