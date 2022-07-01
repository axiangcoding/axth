package axth

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"testing"
	"time"
)

var e *Enforcer

// setup
func setup() {
	db, err := gorm.Open(mysql.Open("root:example@tcp(10.103.4.237:3306)/axth?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
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

func TestEnforcer_Register(t *testing.T) {
	register, err := e.Register(RegisterUser{
		UserID:      "test-user-1",
		DisplayName: "test-user-1",
		Email:       "test-user-1@test.com",
		Phone:       "+8600000000001",
		Password:    "abc",
	})
	if err != nil {
		t.Error(err)
	}
	if !register {
		t.Fail()
	}
}

func TestEnforcer_Login(t *testing.T) {
	_, err := e.LoginWithEmail("test-user-1@test.com", "abc")
	if err != nil {
		t.Error(err)
	}
}

func TestEnforcer_ResetPassword(t *testing.T) {
	_, err := e.ResetPassword("test-user-1", "abc", "newPwd")
	if err != nil {
		t.Error(err)
	}
}

func TestEnforcer_FindUser(t *testing.T) {
	user, err := e.FindUser("test-user-1")
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}
