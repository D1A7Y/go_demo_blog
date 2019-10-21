package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//MD5转换
func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//注册
func Register(username string, password string) string {
	o := orm.NewOrm()
	u := User{}
	u.Username = username
	err := o.Read(&u, "Username")
	if err == orm.ErrNoRows {
		u.PasswordHash = md5V(password)
		_, err = o.Insert(&u)
		if err != nil {
			logs.Info(err)
		}
		return ""
	}
	return "用户已经在 "
}

//登录
func Login(username string, password string) (User, string) {
	o := orm.NewOrm()
	u := User{}
	u.Username = username
	e := o.Read(&u, "Username")
	if e == nil {
		if u.PasswordHash == md5V(password) {
			return u, ""
		}
		return User{Username: username}, "密码错误"
	}
	return u, "找不到用户"
}
