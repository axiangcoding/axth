package axth

import (
	"fmt"
	"os"
	"testing"
)

var e *Enforcer

// setup
func setup() {
	defaultOpt, err := DefaultOptions("axth:pwd@tcp(127.0.0.1:3306)/axth?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	enforcer, err := NewEnforcer(defaultOpt)
	if err != nil {
		panic(err)
	}
	e = enforcer
	fmt.Printf("clean table ax_users for testing")
	err = e.db.Unscoped().Where("1=1").Delete(&AxthUser{}).Error
	if err != nil {
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
