package main

import "fmt"

//嵌入结构，组合
type A struct {
	//嵌入 B 结构
	B
	//嵌入 C 结构
	C //当Name 字段不在本结构定义，并且 Name 在 B 与 C 结构中都有出现，那么无法直接使用 A.Name 写法，否则编译出错
	//Name string
}

type B struct {
	Name string
}

type C struct {
	//Name string
}

func main() {
	var a A = A{}
	a.B.Name = "b"
	//a.C.Name = "c";

	//当Name 字段不在本结构定义，并且 Name 在 B 与 C 结构中都有出现，那么无法直接使用 A.Name 写法，否则编译出错
	fmt.Println(a.Name)
	fmt.Println(a.B.Name)
	fmt.Println("Hello World!")
}
