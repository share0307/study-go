package main

import "fmt"

//嵌入结构，组合
type A struct {
	//嵌入 B 结构
	B
	Name string
}

type B struct {
	Name string
}

func main() {
	var a A = A{Name: "kkk", B: B{Name: string(65)}}

	fmt.Println(a.Name, a.B.Name)
	fmt.Println("Hello World!")
}
