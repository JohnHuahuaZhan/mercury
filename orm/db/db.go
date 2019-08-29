package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

//初始化DB链接
func Init(driverName, dns string) error {
	var err error
	db, err = sqlx.Open(driverName, dns)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return nil
}
func GetDB() *sqlx.DB {
	return db
}
