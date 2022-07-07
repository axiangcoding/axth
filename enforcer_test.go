package axth

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"strconv"
	"testing"
	"time"
)

var e *Enforcer

var tests = []struct {
	user RegisterUser
}{
	{user: RegisterUser{
		UserID:      "dbbc4ba3-c5bb-f559-4c55-dd884c58fb41",
		DisplayName: "test-user-1",
		Email:       "test-user-1@test.com",
		Phone:       "15000000001",
		Password:    "abc",
	}},
	{user: RegisterUser{
		UserID:      "27d86f82-e0a2-d299-c1b3-04ccbd6f0398",
		DisplayName: "test-user-2",
		Email:       "test-user-2@test.com",
		Phone:       "17800000001",
		Password:    "A.bc1234",
	}},
	{user: RegisterUser{
		UserID:      "d3f0c8ac-04f0-a7a0-8ee8-5c0da9e5bb1c",
		DisplayName: "test-user-3",
		Email:       "test-user-3@test.com",
		Phone:       "18800000001",
		Password:    "I'mP@ssword.",
	}},
}

func TestEnforcer_Register(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			register, err := e.Register(tt.user)
			if err != nil {
				t.Error()
			}
			if !register {
				t.Fail()
			}
		})
	}
}

func TestEnforcer_Login(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if _, err := e.Login(tt.user.UserID, tt.user.Password); err != nil {
				t.Fail()
			}
			if _, err := e.LoginWithEmail(tt.user.Email, tt.user.Password); err != nil {
				t.Fail()
			}
			if _, err := e.LoginWithPhone(tt.user.Phone, tt.user.Password); err != nil {
				t.Fail()
			}
			// user not exist case
			if _, err := e.Login(tt.user.UserID+"random", tt.user.Password); err == nil || !errors.Is(err, ErrUserNotExist) {
				t.Fail()
			}
			// user password not matched case
			if _, err := e.Login(tt.user.UserID, tt.user.Password+"random"); err == nil || !errors.Is(err, ErrUserPasswordNotMatched) {
				t.Fail()
			}
		})
	}

}

func TestEnforcer_ResetPassword(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			newPwd := "It'sANewPwd."
			if reset, err := e.ResetPassword(tt.user.UserID, tt.user.Password, newPwd); err != nil || !reset {
				t.Fail()
			}
			if _, err := e.Login(tt.user.UserID, newPwd); err != nil {
				t.Fail()
			}
		})
	}
}

func TestEnforcer_FindUser(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			user, err := e.FindUser(tt.user.UserID)
			if err != nil {
				t.Fail()
			}
			assert.Equal(t, tt.user.UserID, user.UserID)
			assert.Equal(t, tt.user.Phone, user.Phone)
			assert.Equal(t, tt.user.DisplayName, user.DisplayName)
			assert.Equal(t, tt.user.Email, user.Email)
			assert.Equal(t, tt.user.AvatarUrl, user.AvatarUrl)
		})
	}
}

func TestEnforcer_UpdateUser(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			avatarUrl := "http://localhost/avatar/1.png"
			if updated, err := e.UpdateUser(tt.user.UserID, AxthUser{
				Status:    UserStatusBanned,
				AvatarUrl: avatarUrl,
			}); err != nil || !updated {
				t.Fail()
			}
			user, err := e.FindUser(tt.user.UserID)
			if err != nil {
				t.Fail()
			}
			assert.Equal(t, UserStatusBanned, user.Status)
			assert.Equal(t, avatarUrl, user.AvatarUrl)
		})
	}
}

func TestEnforcer_CheckExist(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if exist, err := e.CheckUserIdExist(tt.user.UserID); err != nil || !exist {
				t.Fail()
			}

			if exist, err := e.CheckPhoneExist(tt.user.Phone); err != nil || !exist {
				t.Fail()
			}

			if exist, err := e.CheckEmailExist(tt.user.Email); err != nil || !exist {
				t.Fail()
			}
		})
	}
}

// setup
func setup() {
	dsn := "axth:pwd@tcp(127.0.0.1:3306)/axth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{SingularTable: true}})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic(err)
	}
	defaultOpt, err := DefaultOptions()
	if err != nil {
		panic(err)
	}
	enforcer, err := NewEnforcer(db, defaultOpt)
	if err != nil {
		panic(err)
	}
	e = enforcer
	if err := e.db.Unscoped().Where("1=1").Delete(&AxthUser{}).Error; err != nil {
		panic(err)
	}
}

// teardown
func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
