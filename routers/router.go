package routers

import (
	"demo1/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

func init() {
	//拦截器
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)

	//路由
	//beego.Router("/login", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")

	beego.Router("/user/*", &controllers.UserController{})

}

func TransparentStatic(c *context.Context) {
	path := c.Request.URL.Path // /login.html => /static/login.html

	if strings.Index(path, "user") > 0 {
		return
	}

	http.ServeFile(c.ResponseWriter, c.Request, "static"+path)
}
