package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func Release(title string, content string, author string, time string) string {
	o := orm.NewOrm()
	article := Article{}
	article.Title = title
	article.Content = content
	article.Author = author
	article.Time = time
	_, err := o.Insert(&article)
	if err != nil {
		logs.Info(err)
		return "err"
	}
	return "发布成功"

}

func LookOne(id int) Article {
	o := orm.NewOrm()
	article := Article{}
	article.Id = id
	err := o.Read(&article, "Id")
	if err != nil {
		return Article{}
	} else {
		return article
	}
}
func LookAll(author string) []Article {
	o := orm.NewOrm()
	article := []Article{}
	o.QueryTable("article").Filter("author", author).All(&article)
	return article
}
func Look() []Article {
	o := orm.NewOrm()
	article := []Article{}
	o.QueryTable("article").All(&article)
	return article
}
