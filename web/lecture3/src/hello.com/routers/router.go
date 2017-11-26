package routers

import (
	"hello.com/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //https://beego.me/docs/mvc/controller/router.md
    //首页
    beego.Router("/", &controllers.MainController{})
    //博客首页
    beego.Router("/blog", &controllers.BlogController{})
    //分类
    beego.Router("/category",&controllers.CategoryController{})
    //文章
    beego.Router("/topic",&controllers.TopicController{})
    //自定义方法及 RESTful 规则
    //beego.Router("/topic/add",&controllers.TopicController{},"*:Add")
    //自定义方法及 RESTful 规则
    //beego.Router("/topic/view",&controllers.TopicController{},"*:View")
    //自动化路由
    beego.AutoRouter(&controllers.TopicController{})

    //自动匹配
    beego.AutoRouter(&controllers.UserController{})
    //登录
    //beego.Router("/user", &controllers.LoginController{})
    //自动匹配
    beego.AutoRouter(&controllers.TestController{})
}
