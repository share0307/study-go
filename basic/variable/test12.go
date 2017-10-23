package main

import (
	"fmt"
	//"sort"
)

func main() {
	a := [...]int{1, 2435, 679, 222, 3, 0}

	fmt.Println(a)
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)

	fmt.Println("Hello world!")
}
