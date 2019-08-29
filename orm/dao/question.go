package dao

import "github.com/pingguoxueyuan/gostudy/mercury/orm/model"

func CreateQuestion(question *model.Question) (err error) {

	sqlstr := `insert into question(
				 question_id,  caption, content, author_id, category_id)
			   values(?,?,?,?,?)`

	_, err = DB.Exec(sqlstr, question.QuestionId, question.Caption,
		question.Content, question.AuthorId, question.CategoryId)
	if err != nil {
		logger.Error("create question failed, question:%#v, err:%v", question, err)
		return
	}

	return
}
