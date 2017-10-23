package main

import (
	"fmt"
	//"sort"
)

func main() {
	//切片应用的数组的值改变，也会引起切片值的改变
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}

	sa := a[2:5]
	fmt.Println(string(sa))

	sb := a[3:5]
	fmt.Println(string(sb))

	fmt.Println("Hello world!")
}
