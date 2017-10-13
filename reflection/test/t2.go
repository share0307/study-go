package main

import (
    "fmt"
    //https://studygolang.com/articles/8742
    "reflect"
    "io"
)


func main()  {

    var x uint8 = 'x'

    //类型断言
    var r io.Reader
    var w io.Writer
    w = r.(io.Writer|nil)
    fmt.Println(w);

    v := reflect.ValueOf(x)

    //获取 x 的类型
    fmt.Println("type : ",v.Kind())
    //判断 x 是否 reflect.Uint8 类型
    fmt.Println("type is Uint8:",v.Kind() == reflect.Uint8);

    fmt.Println("Hello World!");
}