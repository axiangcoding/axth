package axth

type IEnforcer interface {
	Login(userId string, password string) (*DisplayUser, error)
	LoginWithEmail(email string, password string) (*DisplayUser, error)
	LoginWithPhone(phone string, password string) (*DisplayUser, error)
	ResetPassword(userId string, oldPwd string, newPwd string) (bool, error)
	Register(ru RegisterUser) (bool, error)
	UpdateUser(userId string, user AxthUser) (bool, error)
	FindUser(userId string) (*DisplayUser, error)
	CheckUserIdExist(userId string) (bool, error)
	CheckEmailExist(email string) (bool, error)
	CheckPhoneExist(phone string) (bool, error)
}
