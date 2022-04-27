package axth

import (
	"testing"
)

func TestEnforcer_NewEnforcer(t *testing.T) {
	_, err := NewEnforcer(&Config{
		Dsn: "root:example@tcp(10.103.4.237:3306)/axth?charset=utf8mb4&parseTime=True&loc=Local"})
	if err != nil {
		t.Fatal("init a new enforcer failed")
	}
}
