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

// Login handles user login and JWT generation
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string 
		Password string
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
		http.Error(w, "Invalid email or password !!", http.StatusUnauthorized)
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
		"message": "Login successful\n",
		"token":   tokenString,
	})
}

type SignupRequest struct {
	FirstName    string
	LastName     string
	Email        string
	Password     string
	Branch       string
	BatchYear    string
	EnrollmentNo string
	OTP          string
}

// Signup handles user signup
func Signup(w http.ResponseWriter, r *http.Request) {

	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	var user models.SocietyUser
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err == nil {
		if user.Verified {
			http.Error(w, "User already exists and is verified", http.StatusBadRequest)
			return
		}
		if req.OTP != "" {
			if req.OTP == user.OTP {
				if time.Now().After(user.ExpiresAt) {
					http.Error(w, "OTP expired", http.StatusBadRequest)
					return
				}
				user.Verified = true
				user.OTP = ""
				database.DB.Save(&user)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{"message": "User verified successfully"})
				return
			} else {
				http.Error(w, "Invalid OTP", http.StatusBadRequest)
				return
			}
		}
	}

	if req.OTP == "" {
		// Ensure that the password is not empty
		if req.Password == "" {
			http.Error(w, "Password cannot be empty", http.StatusBadRequest)
			return
		}
		otp, err := helper.GenerateOTP(6)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newUser := models.SocietyUser{
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			Branch:       req.Branch,
			BatchYear:    req.BatchYear,
			Email:        req.Email,
			EnrollmentNo: req.EnrollmentNo,
			Verified:     false,
			OTP:          otp,
			ExpiresAt:    time.Now().Add(5 * time.Minute),
		}
		// Hash the password before saving it to the database
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newUser.Password = string(hashedPassword)
		if err := database.DB.Create(&newUser).Error; err != nil {
			http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
			return
		}

		emailBody := fmt.Sprintf(`<p>Dear User,</p>
        <p>Welcome to the BPIT Society Management Website!</p>
        <p>To complete your registration, please use the following One-Time Password (OTP):</p>
        <h2>%s</h2>
        <p>This OTP is valid for the next 10 minutes. Please do not share this code with anyone.</p>
        <p>If you did not request this registration, please ignore this email.</p>
        <p>Thank you for joining our community!</p>
        <p>Best regards,</p>
        <p>BPIT Society Portal Team</p>
        <hr>
        <p>Bhagwan Parshuram Institute of Technology</p>
        <p>College Society Portal/p>
        <p><a href="https://alumni.bpitindia.com/">BPIT Society Website</a></p>`, otp)
		// Send OTP via mail
		if err := helper.SendEmail(req.Email, "OTP for Society Management Website SignUp", emailBody); err != nil {
			http.Error(w, "Failed To Send Email for OTP", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "OTP sent successfully",
		})
	}

	// http.Error(w, "Invalid request", http.StatusBadRequest)
}
