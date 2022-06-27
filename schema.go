package axth

import (
	"gorm.io/gorm"
	"time"
)

const (
	// AxUserStatusNormal user status is normal
	AxUserStatusNormal = "normal"
	// AxUserStatusBanned user status is banned
	AxUserStatusBanned = "banned"
)

const (
	FieldId     = "id"
	FieldEmail  = "email"
	FieldPhone  = "phone"
	FieldUserId = "userId"
)

// AxthUser user schema
type AxthUser struct {
	gorm.Model
	UserID           string `gorm:"uniqueIndex;size:255"`
	DisplayName      string `gorm:"size:255"`
	AvatarUrl        string `gorm:"size:255"`
	Email            string `gorm:"uniqueIndex;size:255"`
	Phone            string `gorm:"uniqueIndex;size:255"`
	Password         string `gorm:"size:255"`
	Status           string `gorm:"size:255"`
	LoginFailedCount int
	LastLoginTime    time.Time
}

func (r AxthUser) ToDisplayUser() *DisplayUser {
	user := DisplayUser{
		ID:               r.ID,
		UserID:           r.UserID,
		DisplayName:      r.DisplayName,
		AvatarUrl:        r.AvatarUrl,
		Email:            r.Email,
		Phone:            r.Phone,
		Status:           r.Status,
		LoginFailedCount: r.LoginFailedCount,
		CreatedAt:        r.CreatedAt,
		UpdatedAt:        r.UpdatedAt,
		LastLoginTime:    r.LastLoginTime,
	}
	return &user
}

// DisplayUser For user display
type DisplayUser struct {
	ID               uint      `json:"id"`
	UserID           string    `json:"user_id"`
	DisplayName      string    `json:"display_name"`
	AvatarUrl        string    `json:"avatar_url"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	Status           string    `json:"status"`
	LoginFailedCount int       `json:"login_failed_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	LastLoginTime    time.Time `json:"last_login_time"`
}

// RegisterUser For user register
type RegisterUser struct {
	UserID      string `json:"user_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Password    string `json:"password,omitempty"`
}

func (u RegisterUser) ToAxUser() *AxthUser {
	user := AxthUser{
		UserID:           u.UserID,
		DisplayName:      u.DisplayName,
		Email:            u.Email,
		Phone:            u.Phone,
		Password:         u.Password,
		Status:           AxUserStatusNormal,
		LoginFailedCount: 0,
		LastLoginTime:    time.Now(),
	}
	return &user
}
