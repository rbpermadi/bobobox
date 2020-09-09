package config

import (
	"os"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/subosito/gotenv"
)

type BbbxConfig struct {
	DB   *sqlx.DB
	Mock sqlmock.Sqlmock
}

var (
	bbbxConfig *BbbxConfig
	once       sync.Once
)

func GetInstanceWithEnv(env string) *BbbxConfig {
	once.Do(func() {
		gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/rbpermadi/bobobox/.env")
		if env != "" {
			os.Setenv("ENV", env)
		}

		db, mock := NewMySQL()

		bbbxConfig = &BbbxConfig{DB: db, Mock: mock}
	})
	return bbbxConfig
}

func GetInstance() *BbbxConfig {
	env := os.Getenv("ENV")
	return GetInstanceWithEnv(env)
}
