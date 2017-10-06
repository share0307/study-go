package main

import "fmt"

//嵌入结构，组合

//人类
type human struct {
    Sex int
    Name string
}

//教师类
type teacher struct {
    //把其他的结构作为元素填到另一个结构，就叫组合，可以使用该组合(如 human)的所有属性跟方法
    human
    Name string
    Age int
}

//学生类
type student struct {
    //当省略了名字，那么默认会直接使用类型的名字作为名称
    //human human
    human   //此处实际上是省略了名字
    Name string
    Age int
    int
}


func main()  {

    a := &teacher{Name:"kkk",Age:20,human:human{Sex:1}}
    b := &student{Name:"student",Age:10,human:human{Sex:0}}

    //直接使用 Sex 字段
    a.Sex = 10;
    //通过 human 使用 Sex 字段，这两种方法都是可行的
    a.human.Sex = 10;
    //但是当嵌入的结构体(human)中与被嵌入的结构体存在同名的字段，那么 a.human.Sex 跟 a.Sex 是不同的
    a.human.Name = "aaaaaaaaaaaaa"
    fmt.Println(a)
    fmt.Println(b)

    fmt.Println("Hello World!");
}
