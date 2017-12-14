package main

import (
	"fmt"
	"time"
)

func main() {

	var c chan int = make(chan int)

	go Test(c)

	var tmp int= <- c
	fmt.Println(tmp)
}

func Test(c chan int)  {
	fmt.Println(c)
	fmt.Println(time.Now())

	c <- 100
}