package controllers

import (
	"demo1/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type UserController struct {
	beego.Controller
}
type UserMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ResponseBody struct {
	Message string `json:"message"`
}

func (c *UserController) Post() {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/user/login":
		c.loginPost()
		break
	case "/user/register":
		c.registerPost()
		break
	}
}
func (c *UserController) Get() {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/user/loginOut":
		c.loginOut()
		break
	}
}
func (c *UserController) loginPost() {
	body := c.Ctx.Input.RequestBody
	um := UserMessage{}
	err := json.Unmarshal(body, &um)
	rb := ResponseBody{}
	if err != nil {
		logs.Info(err)
		rb.Message = "???"
	}
	u, e := models.Login(um.Username, um.Password)
	if e == "" {
		rb.Message = "登陆成功"
		c.SetSession("username", &u)
		c.Ctx.SetCookie("username", u.Username, time.Second*3600)
	} else if e == "找不到用户" {
		rb.Message = "找不到用户"
	} else if e == "密码错误 " {
		rb.Message = "密码错误"
	} else {
		rb.Message = "登录错误?"
		logs.Info(e)
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}
func (c *UserController) registerPost() {
	body := c.Ctx.Input.RequestBody
	um := UserMessage{}
	err := json.Unmarshal(body, &um)
	rb := ResponseBody{}
	if err != nil {
		logs.Info(err)
	}
	if um.Username != "" && um.Password != "" {
		e := models.Register(um.Username, um.Password)
		if e == "用户已经在 " {
			rb.Message = "用户已存在"
		} else if e == "" {
			rb.Message = "注册成功"
		} else {
			rb.Message = "注册失败"
			logs.Info(err)
		}
	} else if um.Username == "" {
		rb.Message = "账号为空"
	} else {
		rb.Message = "密码为空"
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}
func (c *UserController) loginOut() {
	u := c.GetSession("username")
	rb := ResponseBody{}
	if u != nil {
		c.DelSession("username")
		rb.Message = "退出成功"
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}
