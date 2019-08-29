package dao

import (
	"database/sql"
	"github.com/JohnHuahuaZhan/mercury/orm"
	"github.com/JohnHuahuaZhan/mercury/orm/db"
	"github.com/JohnHuahuaZhan/mercury/orm/model"
)

//question_id,  caption, content, author_id, category_id 必须设置。默认值也可以
func InsertQuestion(question *model.Question) error {

	sqlstr := `insert into question(
				 question_id,  caption, content, author_id, category_id)
			   values(?,?,?,?,?)`

	_, err := db.GetDB().Exec(sqlstr, question.QuestionId, question.Caption,
		question.Content, question.AuthorId, question.CategoryId)
	if nil != err {
		return orm.ErrSql
	}
	return nil
}

//QuestionByQID 不存在则返回nil
func QuestionByQID(questionId uint64) (*model.Question, error) {
	question := &model.Question{}
	sqlstr := `select * from question where question_id=?`

	err := db.GetDB().Get(question, sqlstr, questionId)
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
