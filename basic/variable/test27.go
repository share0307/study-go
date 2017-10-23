package main

import (
	"fmt"
	//"sort"
)

func main() {

	A(10, "aaa")
	B(10, 20, 30)

	//匿名函数
	c := func(a int) (int, int) {
		return a, a
	}

	fmt.Println(c(10))

	fmt.Println("Hello world!")
}

/*
当传参个数位置，可以直接写成  A(a ...int),实际上 a 是一个切片,必须作为最后一个参数
*/
func A(a int, b string) (int, string) {
	fmt.Println(a, b)

	return a, b
}

func B(a ...int) {
	fmt.Println(a)
}
