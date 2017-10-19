package main

import (
    "fmt"
    //https://studygolang.com/articles/8742
    //"reflect"
    "reflect"
)

type User struct {
    Id int
    Name string
    Age int
}


func main()  {

    u := User{Name:"Hello World!",Id:10,Age:100}

    SetUserName(&u)

    fmt.Println("user:",u)

    fmt.Println("Hello World!");
}


func SetUserName(o interface{})  {
    //或者 o 值
    v := reflect.ValueOf(o)
    //
    //t := reflect.TypeOf(o)

    // 判断 v 值是否可以被修改。只有可寻址的 v 值可被修改。
    // 结构体中的非导出字段（通过 Field() 等方法获取的）不能修改，所有方法不能修改。
    //func (v Value) CanSet() bool
    if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){
        //意思是当为指针类型并且不可修改
        fmt.Println("意思是当为指针类型并且不可修改");
        return
    }else{
        //.Elem() 返回的是 refelect.value 类型的值
        //有什么区别呢？
        //v1:是传进来的引用的值
        //v2:是指针所指对象
        //fmt.Println("v1:",&v)
        v = v.Elem()
    }

    //fmt.Println("v2:",&v)

    f := v.FieldByName("Name");
    // 判断 v 本身（不是 v 值）是否为零值。
    // 如果 v 本身是零值，则除了 String 之外的其它所有方法都会 panic。
    if !f.IsValid(){
        panic("字段不存在!");
    }
    if f.Kind() == reflect.String{
        f.SetString("Hello KKK!")
    }

    //fmt.Println("type:",t.Kind())
    //fmt.Println("value:",v.Kind())
}

//
//// 判断 v 本身（不是 v 值）是否为零值。
//// 如果 v 本身是零值，则除了 String 之外的其它所有方法都会 panic。
//
//// 示例
//func main() {
//    var v reflect.Value      // 未包含任何数据
//    fmt.Println(v.IsValid()) // false
//
//    var i *int
//    v = reflect.ValueOf(i)   // 包含一个指针
//    fmt.Println(v.IsValid()) // true
//
//    v = reflect.ValueOf(nil) // 包含一个 nil 指针
//    fmt.Println(v.IsValid()) // false
//
//    v = reflect.ValueOf(0)   // 包含一个 int 数据
//    fmt.Println(v.IsValid()) // true
//}
//
//------------------------------