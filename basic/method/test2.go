package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

/*
   1.(a A)=>(变量 结构)
   2.意思是这个方法的接收者是 结构A
   3.可以通过 变量a 访问或者使用 结构A 的东西
   4.当接收者不同，可以允许相同的方法名称
   5.当接收者不同，方法不成相同，可是接收参数不同，也是不允许的
*/
//指针传递
func (a *A) Print() {
	a.Name = "a:kkk"
	fmt.Println(a.Name)
}

//值传递
func (b B) Print() {
	b.Name = "b:kkk"
	fmt.Println(b.Name)
}

func main() {
	a := A{}
	a.Name = "a"
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Name = "b"
	b.Print()
	fmt.Println(b.Name)
	fmt.Println("Hello World!")
}
