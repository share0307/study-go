package main

import (
	"runtime"
	"fmt"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		//https://studygolang.com/articles/3028
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		fmt.Printf("%d ", i)
	}
	quit <- 0
}


func main() {

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<- quit
	}
}