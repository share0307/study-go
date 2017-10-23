package main

import (
	"fmt"
	//"sort"
)

func main() {
	//make(类型,初始长度,容量)
	//当长度大于容量，那么容量会 double 扩容
	//当容量够用，那么在底层使用的是一个连续的内存块
	s1 := make([]int, 3, 10)

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	fmt.Println("Hello world!")
}
