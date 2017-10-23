package main

import (
	"fmt"
	//"sort"
)

func main() {

	fmt.Println("a")

	//defer 是整个方法 执行完以后执行的，执行的数序是先定义的，后执行
	defer fmt.Println("b")
	defer fmt.Println("c")

	fmt.Println("Hello world!")
}
