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

type Manage struct {
    User User
    title   string
}

func (u User)Hello()  {
    fmt.Println("Hello ",u.Name)
}

func main()  {
    manage := Manage{User:User{Id:10,Name:"kkk",Age:100},title:"i'm title!"}

    //fmt.Println(manage);

    //http://www.cnblogs.com/golove/p/5909541.html
    t := reflect.TypeOf(manage);

    //fmt.Println(t.Field(0));
    //fmt.Println(t.FieldByName("User"))


    //reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x10b7220), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
    //reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x10b7200), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
    fmt.Printf("%#v\n",t.Field(0))
    fmt.Printf("%#v\n",t.Field(0).PkgPath)

    //通过索引取字段
    //[]int{0,0}是一个切片，然后第一位标识的是 “t” 的第一个字段=>a，第二位标识的是 a 的第一个字段,那么也就意味着 a 应该就是一个内嵌的结构体
    //那么以下的操作就是取 Manage.User.Id 了！
    fmt.Printf("FieldByIndex：%#v=>",t.FieldByIndex([]int{0,0}));

    fmt.Println("Hello World!");
}




