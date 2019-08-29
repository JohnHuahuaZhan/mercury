package service

import (
	"github.com/JohnHuahuaZhan/mercury/orm/dao"
	modelDao "github.com/JohnHuahuaZhan/mercury/orm/model"
	"github.com/JohnHuahuaZhan/mercury/service/model"
	"github.com/JohnHuahuaZhan/mercury/unique"
)

//caption, content, author_id, category_id 必须设置为有意义的值。虽然默认值也可以
//会修改传入的qid，
func CreateQuestion(question *model.QuestionInfo) error {
	questionDao := &modelDao.Question{}
	//生成全局唯一USER ID
	qid, err := unique.GetId()
	question.QuestionId = qid
	questionDao.QuestionId = qid
	questionDao.Content = question.Content
	questionDao.Caption = question.Caption
	questionDao.AuthorId = question.AuthorId
	questionDao.CategoryId = question.CategoryId

	c, err := dao.CategoryByCID(questionDao.CategoryId)
	if nil != err {
		return ErrUnKnow
	}
	if nil == c {
		return ErrCategoryNotExist
	}
	err = dao.InsertQuestion(questionDao)
	if nil != err {
		return ErrUnKnow
	}
	return nil
}
