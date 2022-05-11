package axth

import (
	"fmt"
	"github.com/axiangcoding/axth/data/schema"
	errs "github.com/axiangcoding/axth/errors"
	"github.com/axiangcoding/axth/security"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	loginFieldId     = "id"
	loginFieldEmail  = "email"
	loginFieldPhone  = "phone"
	loginFieldUserId = "userId"
)

type Enforcer struct {
	Db      *gorm.DB
	CacheDb *redis.Client
}

type Config struct {
	// Relational database dsn
	DBDsn string
	// Cache database dsn
	CacheDsn string
}

// NewEnforcer create a new enforcer
func NewEnforcer(config *Config) (*Enforcer, error) {
	db, err := gorm.Open(mysql.Open(config.DBDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&schema.AxthUser{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	opt, err := redis.ParseURL(config.CacheDsn)
	if err != nil {
		return nil, err
	}
	redisClient := redis.NewClient(opt)
	return &Enforcer{
		Db:      db,
		CacheDb: redisClient,
	}, nil
}

// Login 使用用户ID登录，默认登录方式
func (e *Enforcer) Login(userID string, password string) (*schema.DisplayUser, error) {
	return e.loginWithKey(loginFieldUserId, userID, password)
}

// LoginWithEmail 使用邮箱登录
func (e *Enforcer) LoginWithEmail(email string, password string) (*schema.DisplayUser, error) {
	return e.loginWithKey(loginFieldEmail, email, password)
}

// LoginWithPhone 使用手机号登录
func (e *Enforcer) LoginWithPhone(phone string, password string) (*schema.DisplayUser, error) {
	return e.loginWithKey(loginFieldPhone, phone, password)
}

// ResetPassword 重置密码
func (e *Enforcer) ResetPassword(userID string, oldPwd string, newPwd string) (bool, error) {
	where := schema.AxthUser{
		UserID: userID,
	}
	var found schema.AxthUser
	err := e.Db.Where(where).Take(&found).Error
	if err != nil {
		return false, err
	}
	err = security.ComparePwd(found.Password, oldPwd)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	newHashPwd, err := security.GeneratePwd(newPwd)
	if err != nil {
		return false, err
	}
	err = e.Db.Model(&found).Updates(schema.AxthUser{Password: newHashPwd}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// Register 注册一个用户
func (e *Enforcer) Register(ru schema.RegisterUser) (bool, error) {
	user := ru.ToAxUser()
	hashedPassword, err := security.GeneratePwd(ru.Password)
	if err != nil {
		return false, err
	}
	user.Password = string(hashedPassword)
	err = e.Db.Save(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// FindUser 查找一个用户
func (e *Enforcer) FindUser(userID string) (schema.DisplayUser, error) {
	// TODO implement me
	panic("implement me")
}

// CheckUserIdExist 检查用户ID是否存在
func (e *Enforcer) CheckUserIdExist(userId string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

// CheckEmailExist 检查邮箱是否存在
func (e *Enforcer) CheckEmailExist(email string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

// CheckPhoneExist 检查手机号是否存在
func (e *Enforcer) CheckPhoneExist(phone string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (e *Enforcer) loginWithKey(key string, val interface{}, password string) (*schema.DisplayUser, error) {
	where := schema.AxthUser{}
	if key == loginFieldEmail {
		where.Email = val.(string)
	} else if key == loginFieldUserId {
		where.UserID = val.(string)
	} else if key == loginFieldPhone {
		where.Phone = val.(string)
	} else {

		return nil, errs.ErrInternalFailed
	}
	var found schema.AxthUser
	err := e.Db.Where(where).Take(&found).Error
	if err != nil {
		return nil, errs.ErrUserNotExist
	}
	err = security.ComparePwd(found.Password, password)
	if err != nil {
		return nil, errs.ErrUserPasswordNotMatched
	}
	return found.ToDisplayUser(), nil
}
