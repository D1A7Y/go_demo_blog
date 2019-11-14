package controllers

import (
	"demo1/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type ArticleController struct {
	beego.Controller
}
type ArticleMessage struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type ArticleMessage2 struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type ArticleMessage3 struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type ArticleMessage4 struct {
	Count     int              `json:"count"`
	Pagecount int              `json:"pagecount"`
	Pageindex int              `json:"pageindex"`
	Article   []models.Article `json:"article"`
}
type ArticleID struct {
	Id int `json:"id"`
}

func (c *ArticleController) Post() {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/article/release":
		c.releasePost()
		break
	case "/article/modify":
		c.modifyPost()
		break
	}
}

func (c *ArticleController) Get() {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/article/lookOne":
		c.lookOne()
		break
	case "/article/lookAll":
		c.lookAll()
		break
	case "/article/look":
		c.look()
		break
	case "/article/delete":
		c.delete()
		break
	}
}
func (c *ArticleController) releasePost() {
	body := c.Ctx.Input.RequestBody
	am := ArticleMessage2{}
	rb := ResponseBody{}
	err := json.Unmarshal(body, &am)
	if err != nil {
		logs.Info(err)
	}
	articleTime := time.Now().Format("2006-01-02 15:04:05")
	e := models.Release(am.Title, am.Content, am.Author, articleTime)
	if e == "err" {
		rb.Message = "err"
	} else if e == "发布成功" {
		rb.Message = "发布成功"
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}

/*func (c *ArticleController) releasePost() {
	body := c.Ctx.Input.RequestBody
	am := ArticleMessage{}
	rb := ResponseBody{}
	err := json.Unmarshal(body, &am)
	if err != nil {
		logs.Info(err)
	}
	author := c.GetSession("username")
	e := models.Release(am.Title, am.Content, author.(string))
	if e == "err" {
		rb.Message = "err"
	} else if e == "发布成功" {
		rb.Message = "发布成功"
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}*/
func (c *ArticleController) modifyPost() {
	body := c.Ctx.Input.RequestBody
	am := ArticleMessage3{}
	rb := ResponseBody{}
	err := json.Unmarshal(body, &am)
	if err != nil {
		logs.Info(err)
	}
	articleTime := time.Now().Format("2006-01-02 15:04:05")
	rb.Message = models.Modify(am.Id, am.Title, am.Content, articleTime)

	c.Data["json"] = &rb
	c.ServeJSON()
}
func (c *ArticleController) lookOne() {
	//查询某个文章
	id, err := c.GetInt("id")
	if err != nil {
		logs.Info(err)
	}
	article := models.LookOne(id)
	c.Data["json"] = &article
	c.ServeJSON()
}
func (c *ArticleController) lookAll() {
	//查询该作者所有文章
	author := c.GetString("author")
	article := models.LookAll(author)
	c.Data["json"] = &article
	c.ServeJSON()
}

/*func (c *ArticleController) lookAll() {
	//查询该作者所有文章
	author := c.GetSession("username").(string)
	article := models.LookAll(author)
	c.Data["json"] = &article
	c.ServeJSON()
}*/

func (c *ArticleController) look() {
	article := models.Look()
	c.Data["json"] = &article
	c.ServeJSON()

}
func (c *ArticleController) delete() {
	rb := ResponseBody{}
	id, err := c.GetInt("id")
	if err != nil {
		logs.Info(err)
	}
	rb.Message = models.Delete(id)
	c.Data["json"] = &rb
	c.ServeJSON()
}
