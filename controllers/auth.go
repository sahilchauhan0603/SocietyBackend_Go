package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	database "github.com/sahilchauhan0603/society/config"
	"github.com/sahilchauhan0603/society/helper"
	models "github.com/sahilchauhan0603/society/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HandleMicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("CLIENT_ID")
	redirectURL := os.Getenv("REDIRECT_URL")
	tenantID := os.Getenv("TENANT_ID")

	log.Printf("CLIENT_ID: %s, REDIRECT_URL: %s, TENANT_ID: %s\n", clientID, redirectURL, tenantID)

	if clientID == "" || redirectURL == "" || tenantID == "" {
		log.Fatal("Missing required environment variables")
	}

	authURL := fmt.Sprintf(
		"https://login.microsoftonline.com/%s/oauth2/v2.0/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=openid profile email",
		tenantID, clientID, redirectURL,
	)

	http.Redirect(w, r, authURL, http.StatusFound)
}

func HandleMicrosoftCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// Exchange the code for a token
	token, err := exchangeCodeForToken(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validating the token and extracting user info
	idToken, ok := token["id_token"].(string)
	if !ok {
		http.Error(w, "No id_token in response", http.StatusInternalServerError)
		return
	}

	// storing jwt Token
	jwtToken, err := ValidateTokenAndGenerateJWT(idToken)
	if err != nil {
		http.Error(w, "Failed to validate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"jwtToken": "%s"}`, jwtToken)
}

func exchangeCodeForToken(authCode string) (map[string]interface{}, error) {
	// Define the token endpoint
	tokenURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", os.Getenv("TENANT_ID"))

	// Prepare the data for the POST request
	data := url.Values{}
	data.Set("client_id", os.Getenv("CLIENT_ID"))
	data.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", os.Getenv("REDIRECT_URL"))

	// Make the POST request
	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %s", body)
	}

	// Parse the response JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return result, nil
}

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.SocietyUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure that the password is not empty
	if user.Password == "" {
		http.Error(w, "Password cannot be empty", http.StatusBadRequest)
		return
	}

	// Hash the password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Store the alumni profile in the database
	if result := database.DB.Create(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login handles user login and JWT generation
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.SocietyUser
	if result := database.DB.Where("email = ?", credentials.Email).First(&user); result.Error != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}

	// Set the token in a cookie (optional, you can remove this if not needed)
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 72),
		HttpOnly: true,
	})

	// Send the token in the response body as well
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   tokenString,
	})
}

// Signup handles user signup

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.SocietyUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure that the email and enrollment number are unique
	var existingUser models.SocietyUser
	if result := database.DB.Where("email = ? OR enrollment_no = ?", user.Email, user.EnrollmentNo).First(&existingUser); result.RowsAffected > 0 {
		http.Error(w, "Email or Enrollment Number already exists", http.StatusConflict)
		return
	}

	// Generate OTP
	otp, err := helper.GenerateOTP(6)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	otpEntry := models.SocietyOTP{
		Email:     user.Email,
		Code:      otp,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.SocietyOTP{}) {
		if err := database.DB.AutoMigrate(&models.SocietyOTP{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Check if an OTP already exists for this email
	var existingOTP models.SocietyOTP
	err = database.DB.Where("email = ?", user.Email).First(&existingOTP).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err == nil {
		// OTP exists, update it
		existingOTP.Code = otp
		existingOTP.ExpiresAt = time.Now().Add(10 * time.Minute)
		if result := database.DB.Save(&existingOTP); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// OTP does not exist, create a new one
		if result := database.DB.Create(&otpEntry); result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
	}

	emailBody := fmt.Sprintf("<p>Thank you for registering on <b>Bpit Alumni Website</b>.</br> Your OTP is %s.</br> This OTP will expire at %s.</p>", otpEntry.Code, otpEntry.ExpiresAt)
	// Send OTP via mail
	err = helper.SendEmail(user.Email, "Registration OTP", emailBody)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP sent successfully",
	})
}
