package axth

import (
	"github.com/axiangcoding/axth/data/schema"
)

type IEnforcer interface {
	Login(userID string, password string) (*schema.DisplayUser, error)
	LoginWithEmail(email string, password string) (*schema.DisplayUser, error)
	LoginWithPhone(phone string, password string) (*schema.DisplayUser, error)
	ResetPassword(userID string, oldPwd string, newPwd string) (bool, error)
	Register(ru schema.RegisterUser) (bool, error)
	FindUser(userId string) (schema.DisplayUser, error)
	CheckUserIdExist(userId string) (bool, error)
	CheckEmailExist(email string) (bool, error)
	CheckPhoneExist(phone string) (bool, error)
}
