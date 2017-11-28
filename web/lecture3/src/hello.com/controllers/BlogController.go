package controllers

import (
	"github.com/astaxie/beego"
    "hello.com/models"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	c.Data["Title"] = "首页 - 我的 beego 博客"
	c.Data["IsHome"] = true
	c.TplName = "blog/home.html"

	condition := make(map[string]string)
	condition["sort"] = "-created"
	condition["cid"] = c.Input().Get("cid")
	c.Data["Topics"],_ = models.TopicList(c.Input().Get("title"),condition)

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	ck,err := c.Ctx.Request.Cookie("username")

	c.Data["Categories"],_ = models.AllCategoryList()

	if err == nil{
		c.Data["Username"] = ck.Value
	}
}
