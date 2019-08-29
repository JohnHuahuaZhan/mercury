package account

import (
	"github.com/JohnHuahuaZhan/mercury/controller"
	"github.com/JohnHuahuaZhan/mercury/service"
	"github.com/JohnHuahuaZhan/mercury/service/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册
//用户只需要传用户名，密码，邮箱, 即可。其他使用字段类型默认
func RegisterHandler(c *gin.Context) {
	var userInfo model.UserInfoSecret
	err := c.BindJSON(&userInfo) //把用户提交上来的json数据直接解析到结构体中
	if err != nil {
		controller.ResponseError(c, controller.ErrCodeParameter, "参数错误")
		return
	}
	//校验
	if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 || len(userInfo.Username) == 0 {
		controller.ResponseError(c, controller.ErrCodeParameter, "参数不能为空")
		return
	}

	//注册
	err = service.Register(&userInfo)
	if err == service.ErrUserExist { //用户已经存在
		controller.ResponseError(c, controller.ErrCodeUserExist, "用户已注册")
		return
	}
	if err != nil {
		controller.ResponseError(c, controller.ErrUnKnow, "未知错误")
		return
	}

	//注册成功
	controller.ResponseSuccess(c, "注册成功", nil) //返回成功
}

//Login
//用户只需要传用户名，密码即可。其他使用字段类型默认
func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "/user/login", nil)
}

//Login
//用户只需要传用户名，密码即可。其他使用字段类型默认
func LoginHandler(c *gin.Context) {
	sess := sessions.Default(c)
	var userInfo model.UserInfoSecret
	err := c.BindJSON(&userInfo) //把用户提交上来的json数据直接解析到结构体中
	if err != nil {
		controller.ResponseError(c, controller.ErrCodeParameter, "参数错误")
		return
	}
	//校验
	if len(userInfo.Password) == 0 || len(userInfo.Username) == 0 {
		controller.ResponseError(c, controller.ErrCodeParameter, "参数不能为空")
		return
	}

	//登录
	info, err := service.Login(&userInfo)
	if err == service.ErrUserPasswordWrong {
		controller.ResponseError(c, controller.ErrCodeUserExist, service.ErrUserPasswordWrong.Error())
		return
	}
	if err != nil {
		controller.ResponseError(c, controller.ErrUnKnow, "未知错误")
		return
	}

	//登录成功
	sess.Set(controller.SessionUserIdKey(), info.UserId)
	sess.Save()
	controller.ResponseSuccess(c, "登录成功", nil) //返回成功
}
func UserHomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "/auth/user/home", "你的个人主页")
}
