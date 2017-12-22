package main

import (
	"fmt"
	"time"
)

var quit chan int // 只开一个信道

func main() {
	count := 1
	quit = make(chan int) // 无缓冲


	go func() {
		quit <- 0 // ok, finished
		fmt.Println(0)
		//其实到1的时候已经阻塞了，因为消费端没再消费了，容量满了，阻塞了..
		quit <- 1 // ok, finished
		fmt.Println(1)
		quit <- 2 // ok, finished
		fmt.Println(2)
	}()

	for i := 0; i < count; i++ {
		fmt.Println("quit:",<- quit)
		//<- quit
	}

	time.Sleep(time.Second)
}
