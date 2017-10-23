package main

import (
	"fmt"
	//"sort"
)

func main() {

	//此处定义一个存储方法的slice
	var fs = [5]func(){}

	for i := 0; i < 5; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}

}
