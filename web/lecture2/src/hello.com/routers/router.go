package routers

import (
	"hello.com/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //首页
    beego.Router("/", &controllers.MainController{})
    //博客首页
    beego.Router("/blog", &controllers.BlogController{})
    //分类
    beego.Router("/category",&controllers.CategoryController{})


    //自动匹配
    beego.AutoRouter(&controllers.UserController{})
    //登录
    //beego.Router("/user", &controllers.LoginController{})
    //自动匹配
    beego.AutoRouter(&controllers.TestController{})
}
