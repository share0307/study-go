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
		close(quit)
	}()


	//原因是range不等到信道关闭是不会结束读取的。也就是如果 缓冲信道干涸了，那么range就会阻塞当前goroutine, 所以死锁咯。
	for q := range quit{
		fmt.Println("quit:",q)
		//<- quit
	}

	//time.Sleep(time.Second)
	fmt.Println(time.Now())
}