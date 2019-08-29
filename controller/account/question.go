package account

import (
	"github.com/JohnHuahuaZhan/mercury/controller"
	"github.com/JohnHuahuaZhan/mercury/service"
	"github.com/JohnHuahuaZhan/mercury/service/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//创建问题
//用户只需要传标题，内容，分类id, 即可。其他使用字段类型默认
func CreateQuestionHandler(c *gin.Context) {
	sess := sessions.Default(c)
	var question model.QuestionInfo
	err := c.ShouldBindWith(&question, binding.FormPost) //把用户提交上来的json数据直接解析到结构体中
	if err != nil {
		controller.ResponseError(c, controller.ErrCodeParameter, "参数错误")
		return
	}
	//校验
	if len(question.Caption) == 0 || len(question.Content) == 0 {
		controller.ResponseError(c, controller.ErrCodeParameter, "标题或者内容不能为空")
		return
	}

	uid, _ := controller.SessionUserId(sess)
	question.AuthorId = uid
	//注册
	err = service.CreateQuestion(&question)
	if err == service.ErrCategoryNotExist {
		controller.ResponseError(c, controller.ErrCodeCategoryNotExist, service.ErrCategoryNotExist.Error())
		return
	}
	if err != nil {
		controller.ResponseError(c, controller.ErrUnKnow, "未知错误")
		return
	}

	//创建成功
	controller.ResponseSuccess(c, "创建成功", map[string]uint64{"question_id": question.QuestionId}) //返回成功
}
