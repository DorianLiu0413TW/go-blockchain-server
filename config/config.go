package config

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	
	"go-blockchain-api/log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int64
	DatabaseName string
}

type RelyingPartyConfig struct {
	ID          string
	Origin      string
	ServiceName string
}

type InternalServerConfig struct {
	URL            string
	AuthPrivateKey *ecdsa.PrivateKey
}

type Config struct {
	TimeZone          *time.Location
	AllowOrigins      []string
	HTTPListenAddress string
	HTTPListenPort    int64
	RelyingParty      RelyingPartyConfig
	Database          DatabaseConfig
	AuthPrivateKey    *ecdsa.PrivateKey
}

var Global *Config

func InitConfig() (err error) {
	Global = new(Config)

	//
	timeZone := os.Getenv("TIME_ZONE")
	if err = validation.Validate(&timeZone, validation.Required); err != nil {
		err = fmt.Errorf(`"TIME_ZONE" %w`, err)
		return
	}
	if Global.TimeZone, err = time.LoadLocation(timeZone); err != nil {
		err = fmt.Errorf(`error on parsing "TIME_ZONE": %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "TIME_ZONE", timeZone))

	//
	Global.HTTPListenAddress = os.Getenv("HTTP_LISTEN_ADDR")
	if err = validation.Validate(&Global.HTTPListenAddress, validation.Required, is.IPv4); err != nil {
		err = fmt.Errorf(`"HTTP_LISTEN_ADDR" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "HTTP_LISTEN_ADDR", Global.HTTPListenAddress))

	//
	httpListenPortString := os.Getenv("HTTP_LISTEN_PORT")
	if err = validation.Validate(&httpListenPortString, validation.Required, is.Int); err != nil {
		err = fmt.Errorf(`"HTTP_LISTEN_PORT" %w`, err)
		return
	}
	if Global.HTTPListenPort, err = strconv.ParseInt(httpListenPortString, 10, 64); err != nil {
		err = fmt.Errorf(`error on parsing "HTTP_LISTEN_PORT": %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%d", "HTTP_LISTEN_PORT", Global.HTTPListenPort))

	//
	allowOriginStrings := os.Getenv("ALLOW_ORIGINS")
	Global.AllowOrigins = strings.Split(allowOriginStrings, ",")
	if err = validation.Validate(Global.AllowOrigins, validation.Required, validation.Each(is.URL)); err != nil {
		err = fmt.Errorf(`"ALLOW_ORIGINS" %w`, err)
		return
	}
	log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "ALLOW_ORIGINS", allowOriginStrings))

	//
	// Global.RelyingParty.ID = os.Getenv("RP_ID")
	// if err = validation.Validate(&Global.RelyingParty.ID, validation.Required); err != nil {
	// 	err = fmt.Errorf(`"RP_ID" %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "RP_ID", Global.RelyingParty.ID))

	//
	// Global.RelyingParty.Origin = os.Getenv("RP_ORIGIN")
	// if err = validation.Validate(&Global.RelyingParty.Origin, validation.Required, is.URL); err != nil {
	// 	err = fmt.Errorf(`"RP_ORIGIN" %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "RP_ORIGIN", Global.RelyingParty.Origin))

	//
	// Global.RelyingParty.ServiceName = os.Getenv("RP_SERVICE_NAME")
	// if err = validation.Validate(&Global.RelyingParty.ServiceName, validation.Required); err != nil {
	// 	err = fmt.Errorf(`"RP_SERVICE_NAME" %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "RP_SERVICE_NAME", Global.RelyingParty.ServiceName))

	//
	// Global.Database.Host = os.Getenv("DB_HOST")
	// if err = validation.Validate(&Global.Database.Host, validation.Required, is.Host); err != nil {
	// 	err = fmt.Errorf(`"DB_HOST" %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "DB_HOST", Global.Database.Host))

	//
	// dbPortString := os.Getenv("DB_PORT")
	// if err = validation.Validate(&dbPortString, validation.Required, is.Port); err != nil {
	// 	err = fmt.Errorf(`"DB_PORT" %w`, err)
	// 	return
	// }
	// if Global.Database.Port, err = strconv.ParseInt(dbPortString, 10, 64); err != nil {
	// 	err = fmt.Errorf(`error on parsing "DB_PORT": %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%d", "DB_PORT", Global.Database.Port))

	//
	// Global.Database.Username = os.Getenv("DB_USERNAME")
	// if err = validation.Validate(&Global.Database.Username, validation.Required); err != nil {
	// 	err = fmt.Errorf(`"DB_USERNAME" %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "DB_USERNAME", Global.Database.Username))

	//
	// Global.Database.Password = os.Getenv("DB_PASSWORD")
	// if err = validation.Validate(&Global.Database.Password, validation.Required); err != nil {
	// 	err = fmt.Errorf(`"DB_PASSWORD" %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, "Loaded environment variable DB_PASSWORD=********")

	//
	// Global.Database.DatabaseName = os.Getenv("DB_NAME")
	// if err = validation.Validate(&Global.Database.DatabaseName, validation.Required); err != nil {
	// 	err = fmt.Errorf(`"DB_NAME" %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, fmt.Sprintf("Loaded environment variable %s=%s", "DB_NAME", Global.Database.DatabaseName))

	//
	// privateKeyString := os.Getenv("CLIENT_AUTH_PRIVKEY")
	// if err = validation.Validate(
	// 	&privateKeyString,
	// 	validation.Required,
	// 	validation.Length(64, 64),
	// 	util.IsHex,
	// ); err != nil {
	// 	err = fmt.Errorf(`"CLIENT_AUTH_PRIVKEY" %w`, err)
	// 	return
	// }
	// var privateKeyBytes []byte
	// if privateKeyBytes, err = hex.DecodeString(privateKeyString); err != nil {
	// 	err = fmt.Errorf(`error on parsing "CLIENT_AUTH_PRIVKEY": %w`, err)
	// 	return
	// }

	// if Global.AuthPrivateKey, err = util.ParsePrivateKey(privateKeyBytes); err != nil {
	// 	err = fmt.Errorf(`error on parsing "CLIENT_AUTH_PRIVKEY": %w`, err)
	// 	return
	// }
	// log.Info(log.LabelStartup, "Loaded environment variable CLIENT_AUTH_PRIVKEY=********")

	return
}
