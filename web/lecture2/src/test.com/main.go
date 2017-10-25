package main

//http://blog.csdn.net/lastsweetop/article/details/78185713?locationNum=5&fps=1
//https://studygolang.com/articles/10347
import "github.com/astaxie/beego"

type HomeController struct {
    beego.Controller
}

func (this *HomeController) get()  {
    this.Ctx.WriteString("Hello World!")
}

func main()  {
    beego.Router("/",&HomeController{})

    beego.Run();
}
