package main

import "fmt"

//嵌入结构，组合
type A struct {
	//嵌入 B 结构
	//当嵌入的结构体被命名了，则无法直接使用 本结构.字段 这种方法使用,只能通过 本结构.命名名称.字段 使用
	sb B
	//Name string
}

type B struct {
	Name string
}

func main() {
	//当字面值是属于嵌入结构的，不得直接使用，如以下的{Name:"kkk"}
	//var a A = A{Name:"kkk"};
	var a A = A{}
	//初始化以后，即可直接使用
	//a.Name = "kkk";
	//当嵌入的结构体被命名了，则无法直接使用 本结构.字段 这种方法使用,只能通过 本结构.命名名称.字段 使用
	a.sb.Name = "kkk"

	//fmt.Println(a.Name)
	fmt.Println(a.sb.Name)
	fmt.Println("Hello World!")
}
