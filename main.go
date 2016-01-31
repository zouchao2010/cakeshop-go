package main

import (
	_ "cakeshop-go/models"
	_ "cakeshop-go/routers"
//	 "cakeshop-go/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
//	"io"

)

func init() {
//	var w io.Writer
	orm.Debug = true
//	orm.DebugLog = orm.NewLog(w)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.147.151:3306)/database?charset=utf8", 30, 30)
	fmt.Println("main.init..........")
}
func main() {
	beego.SetStaticPath("/upload", "upload")

//	beego.BConfig.WebConfig.TemplateLeft="{%"
//	beego.BConfig.WebConfig.TemplateRight="%}"

	beego.Run()

	//	o := orm.NewOrm()
	//	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//	qs := o.QueryTable("user")
	//
	//	users := make([]*models.User,0)
	//	_, _ = qs.All(&users)
	//	fmt.Println(users)
	//	for _, user := range users {
	//		fmt.Println(user.Mobile)
	//
	//	}
}

