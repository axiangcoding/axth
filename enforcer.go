package axth

import (
	"github.com/axiangcoding/axth/data/schema"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Enforcer struct {
	RdDb    *gorm.DB
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
		RdDb:    db,
		CacheDb: redisClient,
	}, nil
}

func (e Enforcer) Login(email string, password string) error {
	// TODO implement me
	panic("implement me")
}

func (e Enforcer) Register(email string, password string) error {
	// TODO implement me
	panic("implement me")
}

func (e Enforcer) Logout(email string) error {
	// TODO implement me
	panic("implement me")
}

func (e Enforcer) ResetPassword(email string, oldPwd string, newPwd string) error {
	// TODO implement me
	panic("implement me")
}
