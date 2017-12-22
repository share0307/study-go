package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	//这里Go仍然在使用单核，for死循环占据了单核CPU所有的资源，而main线和say两个goroutine都在一个线程里面， 所以say没有机会执行。解决方案还是两个：
	//所以无法输出5个 world！
	//设置单核执行
	runtime.GOMAXPROCS(1)
	go say("world") //开一个新的Goroutines执行
	for {
	}
}