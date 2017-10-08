package main

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
    user := User{Id:7,Name:"kkk",Age:100}
    //初始化结构体,此处取得的是一个类型的指针
    //user := new(User)
    //user.Id = 7
    //user.Name = "kkk"
    //user.Age = 100

    fmt.Println(user)
    //fmt.Println(*user)


    Info(user)


    fmt.Println("Hello World!");
}

//o 为值传递
func Info(o interface{})  {
    //使用 reflect.TypeOf 获得接口的类型
    t := reflect.TypeOf(o);
    fmt.Println("Type:",t.Name())


    if k := t.Kind();k == reflect.Struct {
        //使用 reflect.ValueOf 获得接口的字段集合
        v := reflect.ValueOf(o)
        fmt.Println("Fields:")
        for i := 0; i < t.NumField(); i++ {
            f := t.Field(i)
            //fmt.Println("field:",f)
            //Interface:将 v 值转换为空接口类型。v 值必须可转换为接口类型。
            val := v.Field(i).Interface();
            fmt.Printf("%6s : %v = %v \n", f.Name, f.Type, val);
            //fmt.Println()
        }

        for i := 0; i < t.NumMethod(); i++ {
            m := t.Method(i)
            fmt.Println("%6s:%v", m.Name, m.Type)
            fmt.Println("method:", m)
        }
    }else{
        fmt.Println("类型错误：",t.Kind())
    }


}