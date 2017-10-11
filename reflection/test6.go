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

    t := reflect.TypeOf(o)



    fmt.Println("type:",t.Kind())
    fmt.Println("value:",v.Kind())
}




