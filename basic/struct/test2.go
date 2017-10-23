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
	//直接获取 person 的指针
	a := &person{
		Name: "kkk",
		Age:  10,
	}
	//当 a 为指针时，直接操作a 的属性值时是可以直接改变属性值，而不需要写成 *a.Name = "test"
	//以下出错
	//*a.Name = "test";
	//以下正确
	a.Name = "test"
	fmt.Println(a)
	//值传递
	A(a)
	fmt.Println(a)
	//应用传递
	B(a)
	fmt.Println(a)
}

//证明结构体是值传递
func A(per *person) {
	per.Age = 20
	fmt.Println("A:", per)
}

// 使用应用传递
func B(per *person) {
	per.Age = 200
	fmt.Println("B:", per)
}
