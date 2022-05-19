package axth

import (
	"github.com/axiangcoding/axth/data/schema"
	"os"
	"testing"
)

var conf = Config{
	DBDsn:    "axth:pwd@tcp(127.0.0.1:3306)/axth?charset=utf8mb4&parseTime=True&loc=Local",
	CacheDsn: "redis://localhost:6379/0",
}

var e *Enforcer

// setup 初始化
func setup() {
	enforcer, err := NewEnforcer(&conf)
	if err != nil {
		panic(err)
	}
	e = enforcer
}

// teardown 退出清理
func teardown() {

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestEnforcer_Register(t *testing.T) {
	register, err := e.Register(schema.RegisterUser{
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
