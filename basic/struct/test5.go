package main

import "fmt"

//当为匿名字段时(不给字段命名)，在声明时需要严格按照顺序赋值(只要其中有一个是匿名字段都需要按照此规则)
type person struct {
    Name string
    int
}

func main()  {

    //以下为另一种方法，正确的
    var a *person = &person{
        "kkk",
        20,
    }

    fmt.Println(a);


    fmt.Println("Hello World!");
}
