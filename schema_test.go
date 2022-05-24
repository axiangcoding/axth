package axth

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

var examples = []struct {
	axUser      AxthUser
	displayUser DisplayUser
}{
	{AxthUser{
		Model: gorm.Model{
			ID: 1,
		},
		UserID:           "user_id_1",
		DisplayName:      "user_display_name_1",
		Email:            "email1@test.com",
		Phone:            "+86112345678900",
		Status:           AxUserStatusNormal,
		LoginFailedCount: 0,
	}, DisplayUser{
		ID:               1,
		UserID:           "user_id_1",
		DisplayName:      "user_display_name_1",
		Email:            "email1@test.com",
		Phone:            "+86112345678900",
		Status:           AxUserStatusNormal,
		LoginFailedCount: 0,
	}},
	{AxthUser{
		Model: gorm.Model{
			ID: 2,
		},
		UserID:           "user_id_2",
		DisplayName:      "user_display_name_2",
		Email:            "email-2@test.com",
		Phone:            "+86112345678900",
		Status:           AxUserStatusBanned,
		LoginFailedCount: 0,
	}, DisplayUser{
		ID:               2,
		UserID:           "user_id_2",
		DisplayName:      "user_display_name_2",
		Email:            "email-2@test.com",
		Phone:            "+86112345678900",
		Status:           AxUserStatusBanned,
		LoginFailedCount: 0,
	}},
	{},
}

func TestAxthUser_ToDisplayUser(t *testing.T) {
	for i, example := range examples {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			to := example.axUser.ToDisplayUser()
			assert.Equal(t, to, &example.displayUser)
		})
	}
}
