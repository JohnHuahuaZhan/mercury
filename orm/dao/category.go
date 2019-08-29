package dao

import (
	"database/sql"
	"github.com/JohnHuahuaZhan/mercury/orm"
	"github.com/JohnHuahuaZhan/mercury/orm/db"
	"github.com/JohnHuahuaZhan/mercury/orm/model"
	"github.com/jmoiron/sqlx"
)

//category_id category_name 必须设置。默认值也可以
func InsertCategory(category *model.Category, tx *sqlx.Tx) error {

	sqlstr := `insert into category(
				 category_id,  category_name)
			   values(?,?)`

	_, err := db.GetDB().Exec(sqlstr, category.CategoryId, category.CategoryName)
	if nil != err {
		return orm.ErrSql
	}
	return nil
}

//CategoryByCID 不存在则返回nil
func CategoryByCID(categoryId uint64) (*model.Category, error) {
	category := &model.Category{}
	sqlstr := `select * from category where category_id=?`

	err := db.GetDB().Get(category, sqlstr, categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //没有错误，也没数据
		} else {
			return nil, orm.ErrSql
		}
	} else {
		return category, nil
	}
}
