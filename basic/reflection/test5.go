package main

import (
    "fmt"
    //https://studygolang.com/articles/8742
    //"reflect"
    "reflect"
)


func main()  {

    var a int = 1;

    //t := reflect.TypeOf(a);

    //
    var v reflect.Value = reflect.ValueOf(&a)

    // 获取元素类型、获取指针所指对象类型，获取接口的动态类型
    //因为 v 在初始化的时候传入的是一个 a 的指针，指针无法直接使用 Set*系列的方法，所以必须线使用 .Elem() 方法取得 &a 的值
    //这时候会修改到原来 a 的值
    v.Elem().SetInt(100)

    fmt.Println(a)

    fmt.Println("Hello World!");
}




