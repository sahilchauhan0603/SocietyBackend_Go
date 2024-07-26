package models

import "time"

type SocietyRole struct {
	RoleID           int64 `gorm:"primaryKey;autoIncrement"`
	RoleType         string
	Rolename         string
	RoleDescription  string
	LastDateToApply  string
	Responsibilities string
	LinkBySociety    string
	SocietyID        int64         `gorm:"not null;index"`
	Users            []SocietyUser `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SocietyUser struct {
	UserID          int64 `gorm:"primaryKey;autoIncrement"`
	Username        string
	Password        string
	Email           string
	RoleID          int64            `gorm:"not null;index"`
	SocietyID       int64            `gorm:"not null;index"`
	StudentProfiles []StudentProfile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SocietyProfile struct {
	SocietyID           uint `gorm:"primaryKey"`
	SocietyType         string
	SocietyName         string
	SocietyHead         string
	DateOfRegistration  time.Time
	SocietyDescription  string
	SImage              string
	Testimonials        []SocietyTestimonial `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocietyCoordinator  []SocietyCoordinator `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Events              []SocietyEvent       `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Achievements        []SocietyAchievement `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentProfiles     []StudentProfile     `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Galleries           []SocietyGallery     `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	News                []SocietyNews        `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocietyUsers        []SocietyUser        `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocietyRoles        []SocietyRole        `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentAchievements []StudentAchievement `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SocietyAchievements []SocietyAchievement `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentMarkings     []StudentMarking     `gorm:"foreignKey:SocietyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	Testimonials         []SocietyTestimonial `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentMarking       []StudentMarking     `gorm:"foreignKey:EnrollmentNo;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SocietyAchievement struct {
	SocietyID            uint `gorm:"not null;index"`
	SocietyAchievementID uint `gorm:"primaryKey;autoIncrement:false"`
	Title                string
	Description          string
	DateAchieved         time.Time
}

type SocietyEvent struct {
	SocietyID     uint `gorm:"not null;index"`
	EventID       uint 
	Title         string
	Description   string
	EventType     string
	ModeOfEvent   string
	Location      string
	LinkToEvent   string
	EventDateTime time.Time
}

type StudentAchievement struct {
	EnrollmentNo  uint `gorm:"not null;index"`
	SocietyID     uint `gorm:"not null;index"`
	AchievementID uint `gorm:"primaryKey;autoIncrement:false"`
	Title         string
	Description   string
	DateAchieved  time.Time
}

type StudentMarking struct {
	EnrollmentNo  uint `gorm:"not null;index"`
	SocietyID     uint `gorm:"not null;index"`
	MarkingID     uint `gorm:"primaryKey;autoIncrement:false"`
	StudentGrades string
}

type SocietyTestimonial struct {
	EnrollmentNo           uint `gorm:"not null;index"`
	TestimonialID          uint `gorm:"primaryKey;autoIncrement:false"`
	SocietyID              uint `gorm:"not null;index"`
	TestimonialDescription string
}

type SocietyCoordinator struct {
	SocietyID          uint `gorm:"not null;index"`
	CoordinatorID      uint `gorm:"primaryKey;autoIncrement:false"`
	CoordinatorDetails string
}

type SocietyGallery struct {
	SocietyID uint `gorm:"not null;index"`
	GalleryID uint `gorm:"primaryKey;autoIncrement:false"`
	Image     string
}

type SocietyNews struct {
	SocietyID   uint `gorm:"not null;index"`
	NewsID      uint `gorm:"primaryKey;autoIncrement:false"`
	Title       string
	Description string
	DateOfNews  time.Time
	Author      string
}
