package routers

import (
	"demo1/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"net/http"
	"strings"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//拦截器
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)

	//路由
	//beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")

	beego.Router("/user/*", &controllers.UserController{})
	beego.Router("/article/*", &controllers.ArticleController{})

}

func TransparentStatic(c *context.Context) {
	path := c.Request.URL.Path // /login.html => /static/login.html

	if strings.Index(path, "user") > 0 {
		return
	}
	if strings.Index(path, "article") > 0 {
		return
	}

	http.ServeFile(c.ResponseWriter, c.Request, "static"+path)
}
