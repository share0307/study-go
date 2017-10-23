package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a)

	fmt.Println(b)

	c := strconv.Itoa(a)

	fmt.Println(c)
}
