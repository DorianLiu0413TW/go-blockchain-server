package database

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	"go-blockchain-api/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const (
	dbTimeout      string = "10s"
	dbWriteTimeout string = "30s"
	dbReadTimeout  string = "30s"
)

// InitDB initializes DB and pings it to make sure it is connectable
func InitDB() (err error) {
	if config.Global == nil {
		panic("NO CONFIGS")
	}

	username := config.Global.Database.Username
	password := config.Global.Database.Password
	host := config.Global.Database.Host
	port := config.Global.Database.Port
	dbName := config.Global.Database.DatabaseName

	address := fmt.Sprintf("%s:%d", host, port)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&parseTime=true&loc=%s&multiStatements=true",
		username,
		password,
		address,
		dbName,
		dbTimeout,
		dbReadTimeout,
		dbWriteTimeout,
		url.QueryEscape(config.Global.TimeZone.String()), // Ensure time.Time parsing
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return err
}

func DB() (*sqlx.DB, error) {
	if db == nil {
		return nil, errors.New("DB is uninitialized")
	}
	return db, nil
}

func ClearTransition(tx *sqlx.Tx) {
	rollbackRet := tx.Rollback()
	if rollbackRet != sql.ErrTxDone && rollbackRet != nil {
		panic(rollbackRet.Error())
	}
}