package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	//a := person{};
	//a.Name = "kkk";
	//a.Age = 10;n{};
	//a.Name = "kkk";
	//a.Age = 10;

	//添加初始化值，一定要以,结尾
	a := person{
		Name: "kkk",
		Age:  10,
	}
	fmt.Println(a)
	//值传递
	A(a)
	fmt.Println(a)
	//应用传递
	B(&a)
	fmt.Println(a)
}

//证明结构体是值传递
func A(per person) {
	per.Age = 20
	fmt.Println("A:", per)
}

// 使用应用传递
func B(per *person) {
	per.Age = 200
	fmt.Println("B:", per)
}
