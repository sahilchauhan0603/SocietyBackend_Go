package models

import "time"

type SocietyRole struct {
	RoleID           int64 `gorm:"primaryKey;autoIncrement"`
	Rolename         string
	LastDateToApply  string
	Responsibilities string
	LinkBySociety    string
	Users            []SocietyUser `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SocietyUser struct {
	UserID          int64 `gorm:"primaryKey;autoIncrement"`
	Username        string
	Password        string
	Email           string
	RoleID          int64            `gorm:"not null;index"`
	StudentProfiles []StudentProfile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SocietyProfile struct {
	SocietyID          uint `gorm:"primaryKey"`
	SocietyType        string
	SocietyName        string
	SocietyHead        string
	DateOfRegistration time.Time
	SocietyDescription string
	SocietyCoordinator []Coordinator        `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Events             []SocietyEvent       `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Achievements       []SocietyAchievement `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentProfiles    []StudentProfile     `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Galleries          []Gallery            `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	News               []News               `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type StudentProfile struct {
	EnrollmentNo         uint `gorm:"primaryKey"`
	UserID               uint `gorm:"not null;index"`
	FirstName            string
	LastName             string
	Branch               string
	BatchYear            int
	MobileNo             string
	Email                string
	ProfilePicture       string
	SocietyID            uint `gorm:"not null;index"`
	SocietyPosition      string
	StudentContributions string
	DomainExpertise      string
	GithubProfile        *string
	LinkedInProfile      *string
	TwitterProfile       *string
	StudentAchievements  []StudentAchievement `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Testimonials         []Testimonial        `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentMarking       StudentMarking       `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SocietyAchievement struct {
	SocietyID            uint `gorm:"not null;index"`
	SocietyAchievementID uint `gorm:"not null;index"`
	Title                string
	Description          string
	DateAchieved         time.Time
}

type SocietyEvent struct {
	SocietyID     uint `gorm:"not null;index"`
	EventID       uint `gorm:"primaryKey;autoIncrement:false"`
	Title         string
	Description   string
	EventType     string
	ModeOfEvent   string
	Location      string
	EventDateTime time.Time
}

type StudentAchievement struct {
	EnrollmentNo  uint `gorm:"not null;index"`
	AchievementID uint `gorm:"primaryKey;autoIncrement:false"`
	Title         string
	Description   string
	DateAchieved  time.Time
}

type StudentMarking struct {
	EnrollmentNo  uint `gorm:"not null;index"`
	MarkingID     uint `gorm:"primaryKey;autoIncrement:false"`
	StudentGrades string
}

type Testimonial struct {
	EnrollmentNo         uint
	TestimonialID          uint `gorm:"primaryKey;autoIncrement:false"`
	TestimonialDescription string
}

type Coordinator struct {
	SocietyID          uint 
	CoordinatorID      uint `gorm:"primaryKey;autoIncrement:false"`
	CoordinatorDetails string
}

type Gallery struct {
	SocietyID uint 
	GalleryID uint `gorm:"primaryKey;autoIncrement:false"`
	Image     string
}

type News struct {
	SocietyID   uint 
	NewsID      uint `gorm:"primaryKey;autoIncrement:false"`
	Title       string
	Description string
	DateOfNews  time.Time
	Author      string
}
