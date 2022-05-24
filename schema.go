package axth

import (
	"gorm.io/gorm"
	"time"
)

const (
	AxUserStatusNormal = "normal"
	// AxUserStatusBanned 封禁状态
	AxUserStatusBanned = "banned"
)

// AxthUser For user save in database
type AxthUser struct {
	gorm.Model
	UserID           string `gorm:"uniqueIndex"`
	DisplayName      string `gorm:"size:255"`
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
	ID               uint
	UserID           string
	DisplayName      string
	Email            string
	Phone            string
	Status           string
	LoginFailedCount int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	LastLoginTime    time.Time
}

// RegisterUser For user register
type RegisterUser struct {
	UserID      string
	DisplayName string
	Email       string
	Phone       string
	Password    string
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
