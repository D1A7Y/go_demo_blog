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
func Modify(id int, title string, content string, time string) string {
	o := orm.NewOrm()
	article := Article{}
	article.Id = id
	err := o.Read(&article, "Id")
	if err != nil {
		logs.Info(err)
		return "err1"
	}
	article.Content = content
	article.Title = title
	article.Time = time
	_, err = o.Update(&article, "Title", "Content", "Time")
	if err != nil {
		logs.Info(err)
		return "err2"
	}
	return "成功"

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
	o.QueryTable("article")
	return article
}

/*func Look() (int64,[]Article) {

	/分页
	o := orm.NewOrm()
	article := []Article{}
	qs:=o.QueryTable("article")
	count,err:=qs.Count()
	if err!=nil {
		logs.Info(err)
		return 0,article
	}

	return count,article

}*/
func Delete(id int) string {
	o := orm.NewOrm()
	article := Article{}
	article.Id = id
	_, err := o.Delete(&article)
	if err != nil {
		return "err"
	}
	return "成功"
}
