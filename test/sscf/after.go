package main

import (
	"time"
	"fmt"
)

func main() {

	var after <-chan time.Time

	fmt.Println(time.Now())

	after = time.After(2 * time.Second)

	a := <- after

	fmt.Println(time.Now())
	fmt.Println(a)

}