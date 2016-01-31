package controllers

import (
	"cakeshop-go/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	uid := c.GetString("uid")
	fmt.Println(uid)
	o := orm.NewOrm()
	ads := make([]*models.Ad, 0)
	shops := make([]*models.Shop, 0)
	newest := make([]*models.Shop, 0)
	recomm := make([]*models.Shop, 0)
	o.QueryTable("ad").Limit(6).All(&ads)
	o.QueryTable("shop").Exclude("cid", 2).Exclude("status",9).OrderBy("-views").Limit(6).All(&shops,"name","ename","cover","price")
	for _, shop := range(shops){
		shop.Price =  strings.Split(shop.Price, "~")[0]
		newest = append(newest, shop)
	}
	shops = shops[:0]
	o.QueryTable("shop").Filter("status",1).Limit(6).All(&shops,"name","ename","cover","price")
	for _, shop := range(shops){
		shop.Price =  strings.Split(shop.Price, "~")[0]
		recomm = append(recomm, shop)
	}
	fmt.Println(len(ads))
	c.Data["ads"] = ads
	c.Data["newest"] = newest
	c.Data["recomm"] = recomm
	c.TplName = "site/index.html"
}
