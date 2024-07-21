package models

import "time"

type User struct {
	UserID          int64            `gorm:"primaryKey;autoIncrement"`
	Username        string           
	Password        string          
	Email           string           
	RoleID          int64            `gorm:"not null;index"`
	StudentProfiles []StudentProfile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Role struct {
	RoleID           int64   `gorm:"primaryKey;autoIncrement"`
	Rolename         string 
	LastDateToApply  string 
	Responsibilities string 
	LinkBySociety    string 
	Users             []User `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SocietyAchievement struct {
	SocietyID            uint      `gorm:"not null;index"`
	SocietyAchievementID uint      `gorm:"not null;index"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	DateAchieved         time.Time `json:"date_achieved"`
}

type SocietyEvent struct {
	SocietyID     uint      `gorm:"not null;index"`
	EventID       uint      `gorm:"primaryKey;autoIncrement:false"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	EventType     string    `json:"event_type"`
	ModeOfEvent   string    `json:"mode_of_event"`
	Location      string    `json:"location"`
	EventDateTime time.Time `json:"event_date_time"`
}

type SocietyProfile struct {
	SocietyID          uint                 `gorm:"primaryKey"`
	SocietyType        string               `json:"society_type"`
	SocietyName        string               `json:"society_name"`
	SocietyHead        string               `json:"society_head"`
	SocietyCoordinator string               `json:"society_coordinator"`
	DateOfRegistration time.Time            `json:"date_of_registration"`
	SocietyDescription string               `json:"society_description"`
	Events             []SocietyEvent       `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Achievements       []SocietyAchievement `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentProfiles    []StudentProfile     `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StudentAchievement struct {
	EnrollmentNo  uint      `json:"enrollment_no."`
	AchievementID uint      `gorm:"primaryKey;autoIncrement:false"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	DateAchieved  time.Time `json:"date_achieved"`
}

type StudentMarking struct {
	EnrollmentNo  uint   `gorm:"primaryKey;autoIncrement:false"`
	StudentGrades string `json:"student_grades"`
}

type StudentProfile struct {
	EnrollmentNo         uint `gorm:"primaryKey"`
	UserID               uint `gorm:"not null;index"`
	FirstName            string
	LastName             string
	Branch               string
	BatchYear            int
	MobileNo             string
	ProfilePicture       string
	SocietyID            uint `gorm:"not null;index"`
	SocietyPosition      string
	StudentContributions string
	StudentAchievements  []StudentAchievement `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Testimonials         []Testimonial        `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentMarking       StudentMarking       `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Testimonial struct {
	EnrollmentNo           uint   `gorm:"not null;index"`
	TestimonialDescription string `json:"testimonial_description"`
}
