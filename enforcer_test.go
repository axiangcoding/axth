package axth

import (
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
	registerUser RegisterUser
}{
	{registerUser: RegisterUser{
		UserID:      "dbbc4ba3-c5bb-f559-4c55-dd884c58fb41",
		DisplayName: "test-registerUser-1",
		Email:       "test-registerUser-1@test.com",
		Phone:       "15000000001",
		Password:    "abc",
	}},
	{registerUser: RegisterUser{
		UserID:      "27d86f82-e0a2-d299-c1b3-04ccbd6f0398",
		DisplayName: "test-registerUser-2",
		Email:       "test-registerUser-2@test.com",
		Phone:       "17800000001",
		Password:    "A.bc1234",
	}},
	{registerUser: RegisterUser{
		UserID:      "d3f0c8ac-04f0-a7a0-8ee8-5c0da9e5bb1c",
		DisplayName: "test-registerUser-3",
		Email:       "test-registerUser-3@test.com",
		Phone:       "18800000001",
		Password:    "I'mP@ssword.",
	}},
}

func TestEnforcer_Register(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			register, err := e.Register(tt.registerUser)
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
			_, err := e.LoginWithEmail(tt.registerUser.Email, tt.registerUser.Password)
			if err != nil {
				t.Fail()
			}
		})
	}
}

func TestEnforcer_ResetPassword(t *testing.T) {
	_, err := e.ResetPassword("dbbc4ba3-c5bb-f559-4c55-dd884c58fb41", "abc", "newPwd")
	if err != nil {
		t.Error(err)
	}
}

func TestEnforcer_FindUser(t *testing.T) {
	user, err := e.FindUser("dbbc4ba3-c5bb-f559-4c55-dd884c58fb41")
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
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
