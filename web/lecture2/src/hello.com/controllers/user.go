package controllers

import (
    "github.com/astaxie/beego"
    //"fmt"
    "fmt"
    "github.com/astaxie/beego/context"
)

type UserController struct {
    beego.Controller
}

func (this *UserController)Signin()  {
    this.TplName = "blog/login.html"
}

func (this *UserController)Login()  {
    //this.Ctx.WriteString(fmt.Sprint(this.Input()))
    fmt.Println(fmt.Sprint(this.Input()))
    username := this.Input().Get("username")
    password := this.Input().Get("password")
    autologin := this.Input().Get("autologin") == "on"

    if beego.AppConfig.String("username") == username &&
        beego.AppConfig.String("password") == password {
            maxAge := 0
            if autologin {
                maxAge = 1 << 31 - 1
            }

            this.Ctx.SetCookie("username",username,maxAge,"/")
            this.Ctx.SetCookie("password",password,maxAge,"/")
        }

    this.Redirect("/blog",301)


    return
}

func checkAccount(ctx *context.Context) bool{
    ck,err := ctx.Request.Cookie("username")

    if err != nil{
        return false
    }

    username := ck.Value

    ck,err = ctx.Request.Cookie("password")

    if err != nil{
        return false
    }

    password := ck.Value

    return username == beego.AppConfig.String("username") && password == beego.AppConfig.String("password")
}
