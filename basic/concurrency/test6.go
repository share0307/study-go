package main

import (
	"fmt"
	"runtime"
	//"time"
)

func main() {
	//runtime.GOMAXPROCS：设置 go routine 运行时所用到的核数
	//runtime.NumCPU()：取得当前 cpu 的核数
	//获取当前 go routine 个数
	fmt.Println(runtime.NumGoroutine())
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(" 这台电脑的 cpu 核数是:", runtime.NumCPU())

	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	// 当关闭 channel 后再次遍历管道，那么将会拿到类型的默认值！
	//close(c)
	for i := 0; i < 10; i++ {
		fmt.Println("v:", <-c)
	}

}

func Go(c chan bool, index int) {
	a := 1

	for i := 0; i < 100000; i++ {
		a += i
	}

	fmt.Println(index, a)

	c <- true
}
