package account

import (
	"github.com/JohnHuahuaZhan/mercury/controller"
	"github.com/JohnHuahuaZhan/mercury/service"
	"github.com/JohnHuahuaZhan/mercury/service/model"
	"github.com/gin-gonic/gin"
)

//创建问题
//用户只需要传标题，内容，分类id, 即可。其他使用字段类型默认
func CategoryQuestionsHandler(c *gin.Context) {
	var category model.BasicCategory
	err := c.BindJSON(&category) //把用户提交上来的json数据直接解析到结构体中
	if err != nil {
		controller.ResponseError(c, controller.ErrCodeParameter, "参数错误")
		return
	}

	qs, err := service.QuestionsListCategory(&category)
	if err != nil {
		controller.ResponseError(c, controller.ErrUnKnow, "未知错误")
		return
	}

	//创建成功
	controller.ResponseSuccess(c, "查询成功", qs) //返回成功
}
