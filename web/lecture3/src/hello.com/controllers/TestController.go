package controllers

import "github.com/astaxie/beego"

type TestController struct {
    beego.Controller
}

func (this *TestController)Demo() {
    this.Ctx.WriteString("Hello World!")
}
