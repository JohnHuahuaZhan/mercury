package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/pingguoxueyuan/gostudy/mercury/controller"
	"github.com/pingguoxueyuan/gostudy/mercury/controller/account"
	"github.com/pingguoxueyuan/gostudy/mercury/orm/db"
	"github.com/pingguoxueyuan/gostudy/mercury/unique"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func initDb() (err error) {
	dns := "root:123456@tcp(192.168.16.2:3306)/mercury?parseTime=true&charset=utf8"
	err = db.Init("mysql", dns)
	if err != nil {
		return
	}
	return
}

func main() {
	router := gin.Default()

	//初始化数据库
	err := initDb()
	if err != nil {
		panic(err)
	}

	//初始化全局id生成器
	uniConf := unique.Config{1}
	err = unique.Init(uniConf)
	if err != nil {
		panic(err)
	}

	initTemplate(router)

	//sessions
	store := memstore.NewStore([]byte("secret"))
	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 120}) //Also set Secure: true if using SSL, you should though
	router.Use(sessions.Sessions("gin_blog_sid", store))

	router.GET("/index.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "/index.html", nil)
	})
	router.POST("/user/register", account.RegisterHandler)
	router.POST("/user/login", account.LoginHandler)
	router.GET("/user/login", account.LoginGetHandler)

	//需要登录
	authorized := router.Group("/auth")
	authorized.Use(controller.LoginRequired())
	authorized.GET("/user/home", account.UserHomeHandler)

	router.Run(":9090")
}

func initTemplate(engine *gin.Engine) {
	tmpl := template.New("").Funcs(template.FuncMap{
		"now": now,
	})
	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".gohtml") {
			var err error
			tmpl, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	}

	if err := filepath.Walk("views", fn); err != nil {
		panic(err)
	}
	engine.SetHTMLTemplate(tmpl)
}

//now returns current timestamp
func now() time.Time {
	return time.Now()
}
