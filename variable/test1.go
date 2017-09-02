//当前程序的包名
package main;

//daodao'ru导入其他非 main 包
import "fmt";

//产量的定义
const PI = 3.14;

//全局变量声明以及赋值
//var name="gopher";        //可用
var name string="gopher";   //可用
//name:="gopher";           //报错，但在方法内部可用。syntax error: non-declaration statement outside function body

//一般类型声明
type newType int;

//结构的什么声明
type gopher struct {
    //
}

//接口的声明
type golang interface {

}

func main()  {
    fmt.Println("hello world!",name);
}
