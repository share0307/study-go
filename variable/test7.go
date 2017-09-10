package main

import "fmt"

func main() {
	fmt.Println(^7)
	fmt.Println(1024 << 10 >> 10)
	fmt.Println("a" + "b")

	a := 0
	if a > 0 && (10/10) > 1 {
		fmt.Println("ok!")
	}

	fmt.Println(1 & 0)
}
