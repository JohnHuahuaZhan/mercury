package dao

import "github.com/jmoiron/sqlx"

type Dao struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewDao(db *sqlx.DB, tx *sqlx.Tx) *Dao {
	return &Dao{db, tx}
}
func (dao *Dao) GetDB() *sqlx.DB {
	return dao.db
}
func (dao *Dao) GetTX() *sqlx.Tx {
	return dao.tx
}
