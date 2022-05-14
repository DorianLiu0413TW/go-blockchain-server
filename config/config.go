package config

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"time"
	
	"go-"
	
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host string
	Port int64
	DatabaseName string
}

type InternalServerConfig struct {
	url string
	AuthPrivateKey *ecdsa.PrivateKey
}

type Config struct {
	TimeZone *time.Location
	AllowOrigins []string
	HTTPListenAddress string
	HTTPListenPort int64
	Database DatabaseConfig
	AuthPrivateKey *ecdsa.PrivateKey
}

var Global *Config

func InitConfig() (err error) {
  Global = new(Config)
	timeZone := os.Getenv("TIME_ZONE")
	if err = validation.Validate(&timeZone, validation.Required); err!=nil {
		err = fmt.Errorf(`"TIME_ZONE" %w`, err)
	}
	return

	if Global.TimeZone, err = time.LoadLocation(timeZone); err != nil{
		err = fmt.Errorf(`error on parsing "TIME_ZONE": %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "TIME_ZONE", timeZone))

}
