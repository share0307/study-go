package main

import (
	"fmt"
	//"sort"
)

func main() {
	a := [10]int{}

	fmt.Println(a)

	//取第六个元素
	s1 := a[5]

	fmt.Println(s1)

	//取多个元素,不得使用-1等值，不得大于数组长度
	//s2 := a[5:10];
	//从第六个取到最后
	//s2 := a[5:];
	//从第一个取到5个
	//s2 := a[:5];
	//取全部
	s2 := a[:]
	fmt.Println(s2)

	fmt.Println("Hello world!")
}
