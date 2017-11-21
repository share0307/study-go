package controllers

import (
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	c.Data["Title"] = "首页 - 我的 beego 博客"
	c.Data["IsHome"] = true
	c.TplName = "blog/home.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	ck,err := c.Ctx.Request.Cookie("username")
	if err == nil{
		c.Data["Username"] = ck.Value
	}
}
