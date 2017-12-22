package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop(id int) { // id: 该goroutine的标号
	//http://blog.csdn.net/skh2015java/article/details/60330875
	runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 10; i++ { //打印10次该goroutine的标号
		fmt.Printf("%d ", id)
	}
	quit <- 0
}

func main() {
	runtime.GOMAXPROCS(2) // 最多同时使用2个核

	for i := 0; i < 3; i++ { //开三个goroutine
		go loop(i)
	}

	for i := 0; i < 3; i++ {
		<- quit
	}
}