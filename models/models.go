package models

type User struct {
	UserID   int64  `gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleID   int64  `json:"role_id"`
}
