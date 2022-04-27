package schema

import (
	"gorm.io/gorm"
	"time"
)

type AxthUser struct {
	gorm.Model
	UserID           string
	DisplayName      string
	Email            string
	Password         string
	Status           string
	LoginFailedCount int
	LastLoginTime    time.Time
}
