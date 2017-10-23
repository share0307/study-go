package main

import "fmt"

func main() {

	//匿名结构
	//只是把结构写在初始化时而已
	//a := struct {
	//    Name string
	//    Age int64
	//}{
	//    Name:"kkk",
	//    Age:20,
	//}

	a := &struct {
		Name string
		Age  int64
	}{}

	a.Name = "kkk"
	a.Age = 20

	fmt.Println(a)

	fmt.Println("Hello World!")
}
