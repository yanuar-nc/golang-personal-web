package mysql

import (
	"fmt"
	"learning/personal/config"

	"github.com/jmoiron/sqlx"
)

func NewConnection(cfg config.Config) *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.DBName))
	if err != nil {
		panic(err)
	}
	return db
}
