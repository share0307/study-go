package main

//http://blog.csdn.net/lastsweetop/article/details/78185713?locationNum=5&fps=1
//https://studygolang.com/articles/10347
import (
    "github.com/astaxie/beego"
    "fmt"
)

type HomeController struct {
    beego.Controller
}

func (this *HomeController) Get()  {
    this.Ctx.WriteString("Hello World!")
    fmt.Println("Hello World!");
}

func main()  {
    beego.Router("/",&HomeController{})

    beego.Info("test!!!")

    beego.Error("Error!")

    beego.SetLevel(beego.LevelError)

    beego.Run();
}
