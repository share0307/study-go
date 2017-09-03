package main

import (
    "fmt"
)

//var sss = '123';

//看起来编辑器会报告警，可是实际上能运行
//在声明常量组时，要是没声明，那么会使用上一个表达式
const (
    a = len("aaa")
    b
    c = len("sssda")
    d
)

func main()  {
    fmt.Println(a,b,c,d)
}