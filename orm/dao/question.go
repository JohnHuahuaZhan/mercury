package dao

import (
	"database/sql"
	"github.com/JohnHuahuaZhan/mercury/orm"
	"github.com/JohnHuahuaZhan/mercury/orm/model"
	"github.com/jmoiron/sqlx"
)

type QuestionDao struct {
	*Dao
}

func NewQuestionDao(db *sqlx.DB, tx *sqlx.Tx) *QuestionDao {
	return &QuestionDao{NewDao(db, tx)}
}

//question_id,  caption, content, author_id, category_id 必须设置。默认值也可以
func (questionDao *QuestionDao) InsertQuestion(question *model.Question) error {

	sqlstr := `insert into question(
				 question_id,  caption, content, author_id, category_id)
			   values(?,?,?,?,?)`
	var err error
	if nil != questionDao.tx {
		_, err = questionDao.tx.Exec(sqlstr, question.QuestionId, question.Caption,
			question.Content, question.AuthorId, question.CategoryId)
	} else {
		if nil == questionDao.db {
			return orm.ErrDBParameterNull
		} else {
			_, err = questionDao.db.Exec(sqlstr, question.QuestionId, question.Caption,
				question.Content, question.AuthorId, question.CategoryId)
		}
	}
	if nil != err {
		return orm.ErrSql
	}
	return nil
}

//QuestionByQID 不存在则返回nil
func (questionDao *QuestionDao) QuestionByQID(questionId uint64) (*model.Question, error) {
	question := &model.Question{}
	sqlstr := `select * from question where question_id=?`
	var err error

	if nil != questionDao.tx {
		err = questionDao.tx.Get(question, sqlstr, questionId)
	} else {
		if nil == questionDao.db {
			return nil, orm.ErrDBParameterNull
		} else {
			err = questionDao.db.Get(question, sqlstr, questionId)
		}
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //没有错误，也没数据
		} else {
			return nil, orm.ErrSql
		}
	} else {
		return question, nil
	}
}

//QuestionByQID 不存在则返回nil
func (questionDao *QuestionDao) QuestionsByCID(categoryId uint64) ([]*model.Question, error) {
	var questions []*model.Question
	sqlstr := `select * from question where category_id=?`
	var err error

	if nil != questionDao.tx {
		err = questionDao.tx.Select(&questions, sqlstr, categoryId)
	} else {
		if nil == questionDao.db {
			return nil, orm.ErrDBParameterNull
		} else {
			err = questionDao.db.Select(&questions, sqlstr, categoryId)
		}
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //没有错误，也没数据
		} else {
			return nil, orm.ErrSql
		}
	} else {
		return questions, nil
	}
}
