package axth

type IEnforcer interface {
	Login(email string, password string) error
	Register(email string, password string) error
	Logout(email string) error
	ResetPassword(email string, oldPwd string, newPwd string) error
}
