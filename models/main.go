//https://beego.me/docs/mvc/model/orm.md
package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	orm.RegisterDataBase("default", "mysql", "root:root@orm_test?charset=utf8")
}
func main() {
	o := orm.NewOrm()
	o.Using("default") //默认使用default，指定数据库

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}

func read() {
	o := orm.NewOrm()
	user := User{Name: "sss"}
	//查询name等于"sss"的数据
	o.Read(&user, "Name")
}
