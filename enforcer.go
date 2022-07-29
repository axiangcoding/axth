package axth

import (
	"errors"
	"github.com/axiangcoding/axth/security"
	"gorm.io/gorm"
)

type Enforcer struct {
	db      *gorm.DB
	options *Options
}

// NewEnforcer create a new enforcer
func NewEnforcer(db *gorm.DB, opt *Options) (*Enforcer, error) {
	if opt.DbAutoMigrate {
		err := db.AutoMigrate(&AxthUser{})
		if err != nil {
			return nil, err
		}
	}
	return &Enforcer{
		db:      db,
		options: opt,
	}, nil
}

// Login user login with userId and password, default login method
func (e *Enforcer) Login(userId string, password string) (*DisplayUser, error) {
	return e.loginWithKey(FieldUserId, userId, password)
}

// LoginWithEmail user login with email
func (e *Enforcer) LoginWithEmail(email string, password string) (*DisplayUser, error) {
	return e.loginWithKey(FieldEmail, email, password)
}

// LoginWithPhone user login with phone
func (e *Enforcer) LoginWithPhone(phone string, password string) (*DisplayUser, error) {
	return e.loginWithKey(FieldPhone, phone, password)
}

// ResetPassword reset account password
func (e *Enforcer) ResetPassword(userId string, oldPwd string, newPwd string) (bool, error) {
	where := AxthUser{
		UserID: userId,
	}
	var found AxthUser
	if err := e.db.Where(where).Take(&found).Error; err != nil {
		return false, err
	}
	if err := security.ComparePwd(found.Password, oldPwd); err != nil {
		return false, err
	}
	newHashPwd, err := security.GeneratePwd(newPwd)
	if err != nil {
		return false, err
	}
	if err := e.db.Model(&found).Updates(AxthUser{Password: newHashPwd}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// Register register a user
func (e *Enforcer) Register(ru RegisterUser) (bool, error) {
	user := ru.ToAxUser()
	hashedPassword, err := security.GeneratePwd(ru.Password)
	if err != nil {
		return false, err
	}
	user.Password = hashedPassword
	if err := e.db.Save(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

// FindUser find a user
func (e *Enforcer) FindUser(userId string) (*DisplayUser, error) {
	where := AxthUser{UserID: userId}
	var found AxthUser
	if err := e.db.Where(where).Take(&found).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotExist
		}
		return nil, err
	}
	return found.ToDisplayUser(), nil
}

// UpdateUser update user by userId
func (e *Enforcer) UpdateUser(userId string, user AxthUser) (bool, error) {
	if result := e.db.Where(&AxthUser{UserID: userId}).Updates(user); result.Error != nil {
		return false, result.Error
	} else {
		if result.RowsAffected != 1 {
			return false, ErrUserUpdateFailed
		}
	}
	return true, nil
}

// CheckUserIdExist check if userId already exist
func (e *Enforcer) CheckUserIdExist(userId string) (bool, error) {
	return e.checkValueExist(FieldUserId, userId)
}

// CheckEmailExist check if email already exist
func (e *Enforcer) CheckEmailExist(email string) (bool, error) {
	return e.checkValueExist(FieldEmail, email)
}

// CheckPhoneExist check if phone already exist
func (e *Enforcer) CheckPhoneExist(phone string) (bool, error) {
	return e.checkValueExist(FieldPhone, phone)
}

func (e *Enforcer) loginWithKey(key string, val interface{}, password string) (*DisplayUser, error) {
	where := AxthUser{}
	if key == FieldEmail {
		where.Email = val.(string)
	} else if key == FieldUserId {
		where.UserID = val.(string)
	} else if key == FieldPhone {
		where.Phone = val.(string)
	}
	var found AxthUser
	if err := e.db.Where(where).Take(&found).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotExist
		} else {
			return nil, err
		}
	}
	err := security.ComparePwd(found.Password, password)
	if err != nil {
		return nil, ErrUserPasswordNotMatched
	}
	return found.ToDisplayUser(), nil
}

func (e *Enforcer) checkValueExist(key string, val interface{}) (bool, error) {
	where := AxthUser{}
	if key == FieldEmail {
		where.Email = val.(string)
	} else if key == FieldUserId {
		where.UserID = val.(string)
	} else if key == FieldPhone {
		where.Phone = val.(string)
	} else {
		return false, ErrInternalFailed
	}
	var count int64
	err := e.db.Model(&where).Where(where).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count >= 1, nil
}
