package main

import (
	"fmt"
	//"sort"
)

func main() {
	//切片应用的数组的值改变，也会引起切片值的改变
	a := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}

	sa := a[2:5]
	fmt.Println(string(sa))

	//有一个需要注意的是，[sa 的开始值:sa 的结束值],而非原切片 a 的
	sb := sa[1:5]
	fmt.Println(string(sb))

	fmt.Println(len(sa), cap(sa))

	fmt.Println("Hello world!")
}
