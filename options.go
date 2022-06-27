package axth

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type Options struct {
	// Relational database dsn
	DbDsn string `validate:"required"`
	// auto migrate inner table
	DbAutoMigrate bool
	// db max idle connections
	DbMaxIdleConns int `validate:"gte=0,lte=100"`
	// db max open connections
	DbMaxOpenConns int `validate:"gte=0,lte=100"`
	// db connection max lifetime
	DbConnMaxLifeTime time.Duration
	// user max login failed times
	UserMaxLoginFailed int `validate:"gte=0"`
	// if user reach max login failed, wait for duration to unlock
	UserLoginFailedUnlockDuration time.Duration
}

func DefaultOptions(dbDsn string) (*Options, error) {
	options := Options{
		DbDsn:                         dbDsn,
		DbAutoMigrate:                 true,
		DbMaxIdleConns:                10,
		DbMaxOpenConns:                100,
		DbConnMaxLifeTime:             time.Hour,
		UserMaxLoginFailed:            3,
		UserLoginFailedUnlockDuration: time.Minute * 5,
	}
	err := CheckOptions(options)
	if err != nil {
		return nil, err
	}
	return &options, nil
}

func CheckOptions(opt Options) error {
	validate := validator.New()
	err := validate.Struct(&opt)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		// from here you can create your own error messages in whatever language you wish
		return err
	}

	return nil
}
