package axth

import (
	"testing"
)

var conf = Config{
	DBDsn:    "axth:pwd@tcp(127.0.0.1:3306)/axth?charset=utf8mb4&parseTime=True&loc=Local",
	CacheDsn: "redis://localhost:6379/0",
}

func TestEnforcer_NewEnforcer(t *testing.T) {
	_, err := NewEnforcer(&conf)
	if err != nil {
		t.Fatal("init a new enforcer failed")
	}
}

func TestEnforcer_Login(t *testing.T) {
	e, _ := NewEnforcer(&conf)
	err := e.Login("userA", "pwdA")
	if err != nil {
		t.Fail()
	}
}

func TestEnforcer_Logout(t *testing.T) {

}
