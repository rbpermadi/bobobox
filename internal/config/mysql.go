package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Init returns connector to Soteria database
func NewMySQL() (*sqlx.DB, sqlmock.Sqlmock) {

	var db *sqlx.DB
	var mock sqlmock.Sqlmock
	var err error

	env := os.Getenv("ENV")

	if env == "test" || env == "" {

		_, mock, err = sqlmock.NewWithDSN("sqlmock_db")

		db, err = sqlx.Open("sqlmock", "sqlmock_db")
		if err != nil {
			panic(err.Error())
		}

	} else {
		dbUsername := os.Getenv("DATABASE_USERNAME")
		dbPassword := os.Getenv("DATABASE_PASSWORD")
		dbHost := os.Getenv("DATABASE_HOST")
		dbName := os.Getenv("DATABASE_NAME")
		dbPort := os.Getenv("DATABASE_PORT")
		if dbPort == "" {
			dbPort = "3306"
		}

		if env == "development" || env == "staging" {
			fmt.Printf("Connecting to [USERNAME]:[PASSWORD]@(%s:%v)/%s?parseTime=true\n", dbHost, dbPort, dbName)
		}

		dataSourceName := fmt.Sprintf("%s:%v@(%s:%v)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
		db, _ = sqlx.Open("mysql", dataSourceName)

		if err := db.Ping(); err != nil {
			panic(err.Error())
		}

		if dp, err := strconv.Atoi(os.Getenv("DATABASE_POOL")); err == nil && dp > 0 {
			db.SetMaxIdleConns(dp)
		}

		db.SetConnMaxLifetime(time.Minute)
	}

	return db, mock
}
