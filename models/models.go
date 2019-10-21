package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

func init() {
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/demo1?charset=utf8")

	//映射model数据
	orm.RegisterModel(new(User))
	//生成表
	orm.RunSyncdb("default", false, true)
}
