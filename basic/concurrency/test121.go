package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 100; i++ { //为了观察，跑多些
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func main() {
	//http://blog.csdn.net/skh2015java/article/details/60330875
	//如果当前goroutine不发生阻塞，它是不会让出CPU给其他goroutine的, 所以例子一中的输出会是一个一个goroutine进行的，而sleep函数则阻塞掉了 当前goroutine, 当前goroutine主动让其他goroutine执行, 所以形成了逻辑上的并行, 也就是并发。
	//runtime.GOMAXPROCS(runtime.NumCPU()) // 最多使用2个核
	runtime.GOMAXPROCS(1) // 最多使用2个核

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<- quit
	}
}