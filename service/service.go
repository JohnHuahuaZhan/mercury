package service

import (
	"github.com/JohnHuahuaZhan/mercury/orm/db"
	"github.com/jmoiron/sqlx"
)

func DT(tx bool) (*sqlx.DB, *sqlx.Tx) {
	if tx {
		t, _ := db.GetDB().Beginx()
		return nil, t
	} else {
		return db.GetDB(), nil
	}
}
