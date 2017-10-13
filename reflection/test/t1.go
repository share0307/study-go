package test

import (
    "fmt"
    //https://studygolang.com/articles/8742
    "reflect"
)

type User struct {
    Id int
    Name string
    Age int
}

func (u User)Hello()  {
    fmt.Println("Hello ",u.Name)
}

func main()  {

    //初始化结构体,此处取得的是一个struct
    //初始化结构体,此处取得的是一个类型的指针
    //user := new(User)
    //user.Id = 7
    //user.Name = "kkk"
    //user.Age = 100


    var x float64 = 3.4
    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())
    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
    fmt.Println("value:", v.Float())


    fmt.Println("Hello World!");
}