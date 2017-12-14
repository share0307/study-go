package main

import (
	"fmt"
	"time"
)

func main() {

	go a()
	go a()
	go a()


	time.Sleep(2 * time.Second)
}

func a()  {
	var i int=10
	fmt.Println(i)
	fmt.Println(&i)
	i++
}