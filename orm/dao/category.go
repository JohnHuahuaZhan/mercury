package dao

import (
	"database/sql"
	"github.com/JohnHuahuaZhan/mercury/orm"
	"github.com/JohnHuahuaZhan/mercury/orm/model"
	"github.com/jmoiron/sqlx"
)

type CategoryDao struct {
	*Dao
}

func NewCategoryDao(db *sqlx.DB, tx *sqlx.Tx) *CategoryDao {
	return &CategoryDao{NewDao(db, tx)}
}

//category_id category_name 必须设置。默认值也可以
func (categoryDao *CategoryDao) InsertCategory(category *model.Category, tx *sqlx.Tx) error {

	sqlstr := `insert into category(
				 category_id,  category_name)
			   values(?,?)`
	var err error
	if nil != categoryDao.tx {
		_, err = categoryDao.tx.Exec(sqlstr, category.CategoryId, category.CategoryName)
	} else {
		if nil == categoryDao.db {
			return orm.ErrDBParameterNull
		} else {
			_, err = categoryDao.db.Exec(sqlstr, category.CategoryId, category.CategoryName)
		}
	}
	if nil != err {
		return orm.ErrSql
	}
	return nil
}

//CategoryByCID 不存在则返回nil
func (categoryDao *CategoryDao) CategoryByCID(categoryId uint64) (*model.Category, error) {
	category := &model.Category{}
	sqlstr := `select * from category where category_id=?`
	var err error
	if nil != categoryDao.tx {
		err = categoryDao.tx.Get(category, sqlstr, categoryId)
	} else {
		if nil == categoryDao.db {
			return nil, orm.ErrDBParameterNull
		} else {
			err = categoryDao.db.Get(category, sqlstr, categoryId)
		}
	}
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
