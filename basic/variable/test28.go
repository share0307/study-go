package main

import (
	"fmt"
	//"sort"
)

func main() {

	f := closure(10)

	fmt.Println(f(20))

	fmt.Println("Hello world!")
}

func closure(a int) func(int) int {
	return func(y int) int {
		return a * y
	}
}
