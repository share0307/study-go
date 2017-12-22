package main

import (
	"fmt"
	"time"
)

var quit chan int // 只开一个信道

func foo(id int) {
	//fmt.Println(id)
	quit <- id // ok, finished
}

func main() {
	count := 1000
	quit = make(chan int,count) // 无缓冲

	go func() {
		for i := 0; i < count; i++ {
			//go foo(i)
			quit <- i
		}
	}()

	for i := 0; i < count; i++ {
		fmt.Println("quit:",<- quit)
		//<- quit
	}

	//time.Sleep(time.Second)
	fmt.Println(time.Now())
}