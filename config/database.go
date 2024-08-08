package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/sahilchauhan0603/society/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnector() {
	// Retrieve environment variables for the database
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbName == "" || dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Missing required database environment variables")
	}
	// First, check if the database exists and create it if it doesn't
	serverDSN := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/postgres?sslmode=verify-full"
	db, err := sql.Open("postgres", serverDSN)
	if err != nil {
		log.Fatal("failed to connect to MySQL server: ", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("failed to create database: ", err)
	}

	// Now, connect to the newly created or existing database
	dsn := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=verify-full"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	
	//Migrate the schema
	if err := DB.AutoMigrate(
		&models.SocietyProfile{},
		&models.SocietyRole{},
		&models.SocietyUser{},
		&models.StudentProfile{},
		&models.SocietyAchievement{},
		&models.SocietyEvent{},
		&models.StudentAchievement{},
		&models.StudentMarking{},
		&models.SocietyTestimonial{},
		&models.SocietyCoordinator{},
		&models.SocietyGallery{},
		&models.SocietyNews{},
		&models.AdminPanelRole{},
	); err != nil {
		log.Fatal(err)
	}
	checkAndCreateDefaultAdmin()
}

func checkAndCreateDefaultAdmin() {
	var admin models.AdminPanelRole
	result := DB.Where("role = ?", "admin").First(&admin)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// No admin found, create a default admin
		defaultAdmin := models.AdminPanelRole{
			Username: os.Getenv("ADMIN_USER"),
			Password: hashPassword(os.Getenv("ADMIN_PASS")),
			Role:     "admin",
			SocietyID: 0,
		}
		if err := DB.Create(&defaultAdmin).Error; err != nil {
			log.Fatal("failed to create default admin: ", err)
		}
		log.Println("Default admin user created")
	} else if result.Error != nil {
		log.Fatal("failed to check admin existence: ", result.Error)
	} else {
		log.Println("Admin user already exists")
	}
}

// hashPassword hashes a plain text password using bcrypt
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to hash password: ", err)
	}
	return string(hashedPassword)
}

// func DatabaseConnector() {
// 	// Retrieve environment variables for the database
// 	dbName := os.Getenv("DB_NAME")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")

// 	if dbName == "" || dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" {
// 		log.Fatal("Missing required database environment variables")
// 	}

// 	// First, check if the database exists and create it if it doesn't
// 	serverDSN := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/"
// 	db, err := sql.Open("mysql", serverDSN)
// 	if err != nil {
// 		log.Fatal("failed to connect to MySQL server: ", err)
// 	}
// 	defer db.Close()

// 	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
// 	if err != nil {
// 		log.Fatal("failed to create database: ", err)
// 	}

// 	// Now, connect to the newly created or existing database
// 	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("failed to connect database: ", err)
// 	}

// 	// Migrate the schema
// 	if err := DB.AutoMigrate(
// 		&models.Role{},
// 		&models.User{},
// 		&models.SocietyProfile{},
// 		&models.StudentProfile{},
// 		&models.SocietyAchievement{},
// 		&models.SocietyEvent{},
// 		&models.StudentAchievement{},
// 		&models.StudentMarking{},
// 		&models.Testimonial{},
// 		&models.Coordinator{},
// 		&models.Gallery{},
// 		&models.News{},
// 	); err != nil {
// 		log.Fatal(err)
// 	}
// }
