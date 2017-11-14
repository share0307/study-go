package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//beego 中所有需要注入到变量都通过 c.Data 这个 map 变量注入到模板中
	c.Data["Website"] = "test hello.com"
	c.Data["Email"] = "530090596@qq.com"
	c.Data["Hello"] = "<div>Hello World!</div>"
	c.Data["Slice"] = []int{1,2,3,4,5,6,7,8,9,10}
	c.Data["Pipe"] = "<div>Hello World!</div>"

	beego.Error("到这里了吗？！")

	//beego 使用 c.TplName 指定使用的模板
	c.TplName = "index.tpl"
}
