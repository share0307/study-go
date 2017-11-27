package controllers

import (
    "github.com/astaxie/beego"
    "strings"
    "hello.com/models"
)

type CategoryController struct {
    beego.Controller
}

func (this *CategoryController)Get()  {

    op := this.Input().Get("op")

    switch op {
    case "add":
        name := this.Input().Get("name")
        if len(strings.TrimSpace(name)) == 0 {
            break
        }
        //添加分类
        models.AddCategory(name)
        this.Redirect("/category",302)
        return
    case "del":
        id := this.Input().Get("id")
        if len(strings.TrimSpace(id)) == 0{
            break
        }
        //删除分类
        models.DelCategory(id)
        this.Redirect("/category",302)
    default:
        categories,_ := models.AllCategoryList()
        this.Data["IsCategory"] = true
        this.Data["Categories"] = categories
        this.TplName = "category/home.html"
    }

    return
}
