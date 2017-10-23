//当前程序的包名
package main

//daodao'ru导入其他非 main 包
//import [别名] 包名
//但是当别名为 . 时，那么可以直接使用里边的方法
//如 import . "fmt" 那么 fmt.PrintLn(..) 可以写成  Println()
//import fmt "fmt";
//import . "fmt";
import "fmt"

//引入多个包
import (
//"time"
//os "os"
)

//产量的定义
const PI = 3.14

//产量科小写
const test = "test"

//全局变量声明以及赋值
//var name="gopher";        //可用
var name string = "gopher" //可用
//name:="gopher";           //报错，但在方法内部可用。syntax error: non-declaration statement outside function body

//一般类型声明
type newType int

//结构的什么声明
type gopher struct {
	//
}

//接口的声明
type golang interface {
}

func main() {
	fmt.Println("hello world!", name)
}
