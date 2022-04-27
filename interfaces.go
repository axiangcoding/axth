package axth

import "time"

type IEnforcer interface {
	Login(userID string, password string) error
	RefreshToken(userID string, expireTime time.Duration) error
	BannedUser(userID string) error
	Register(userID string, password string) error
	Logout(userID string) error
	ResetPassword(userID string, oldPwd string, newPwd string) error
}
