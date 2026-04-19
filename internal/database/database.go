package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	config "github.com/sahilchauhan0603/society/internal/config"
	"github.com/sahilchauhan0603/society/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnector(cfg *config.Config) error {
	serverDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/postgres?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	if cfg.Database.MaxRetry < 1 {
		cfg.Database.MaxRetry = 1
	}

	db, err := sql.Open("pgx", serverDSN)
	if err != nil {
		return fmt.Errorf("failed to initialize database driver: %w", err)
	}

	var pingErr error
	for attempt := 1; attempt <= cfg.Database.MaxRetry; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		pingErr = db.PingContext(ctx)
		cancel()

		if pingErr == nil {
			break
		}

		if attempt < cfg.Database.MaxRetry {
			log.Printf("database ping failed (attempt %d/%d): %v", attempt, cfg.Database.MaxRetry, pingErr)
			time.Sleep(cfg.Database.RetryGap)
		}
	}

	if pingErr != nil {
		db.Close()
		return fmt.Errorf("failed to connect to database server after retries: %w", pingErr)
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", cfg.Database.Name).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check database existence: %w", err)
	}

	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", quoteIdentifier(cfg.Database.Name)))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
	}

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to application database: %w", err)
	}

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
		return fmt.Errorf("auto-migration failed: %w", err)
	}

	if err := checkAndCreateDefaultAdmin(cfg.Admin.User, cfg.Admin.Password); err != nil {
		return err
	}

	return nil
}

func checkAndCreateDefaultAdmin(username string, password string) error {
	if username == "" || password == "" {
		return errors.New("default admin credentials are required")
	}

	var admin models.AdminPanelRole
	result := DB.Where("role = ?", "admin").First(&admin)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		defaultAdmin := models.AdminPanelRole{
			Username:  username,
			Password:  hashPassword(password),
			Role:      "admin",
			SocietyID: 0,
		}
		if err := DB.Create(&defaultAdmin).Error; err != nil {
			return fmt.Errorf("failed to create default admin: %w", err)
		}
		log.Println("Default admin user created")
	} else if result.Error != nil {
		return fmt.Errorf("failed to check admin existence: %w", result.Error)
	}

	return nil
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to hash password: ", err)
	}
	return string(hashedPassword)
}

func PingContext(ctx context.Context) error {
	if DB == nil {
		return errors.New("database is not initialized")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.PingContext(ctx)
}

func quoteIdentifier(value string) string {
	return `"` + strings.ReplaceAll(value, `"`, `""`) + `"`
}
