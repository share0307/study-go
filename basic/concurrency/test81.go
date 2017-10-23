package main

import (
	"fmt"
)

func main() {

	//用于数据传输
	c1, c2 := make(chan int), make(chan string)
	//用于信号发送
	o := make(chan bool)

	go func() {
		for {
			//http://blog.csdn.net/zhonglinzhang/article/details/45913443
			//此处为了验证在 for 循环(无线循环)中，需要至少从 select 中执行一个才能继续下一个循环
			//每次 for 循环只会执行其中一个 case 或者 default，要是需要监听的队列同时发送数据过来，其实也是等执行完一个再在下一次 select 循环中执行
			select {
			case v, ok := <-c1:
				fmt.Println("c1")
				if !ok {
					//当 c1或者 c2任一有一个关闭的，那么给发送关闭信号
					o <- true
					//其实这个 timeout 在 select 中没什么用
					break
				}
				fmt.Println("c1:", v)
			case v, ok := <-c2:
				fmt.Println("c2")
				if !ok {
					//当 c1或者 c2任一有一个关闭的，那么给发送关闭信号
					o <- true
					break
				}
				fmt.Println("c2:", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "Hi"
	c1 <- 3
	c2 <- "Hello"

	close(c1)
	close(c2)

	oo := <-o
	fmt.Println(oo)

	fmt.Println()
}
