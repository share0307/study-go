package main

import (
	"fmt"
	//"time"
)

func main() {

	//用于数据传输
	c1, c2 := make(chan int), make(chan string)
	//用于信号发送
	o := make(chan bool, 10)

	go func() {
		for {
			//fmt.Println("")
			//当 c1以及 c2都关闭掉，那么在 for 循环里的 select 中会一直收到信号，为的是通知一下该频道被关闭了
			select {
			case v, ok := <-c1:
				if !ok {
					//当 c1或者 c2任一有一个关闭的，那么给发送关闭信号
					o <- true
					fmt.Println("c1  被关闭了！")
					//time.Sleep(1 * time.Second)
					//break 是有效的..
					//break;
				}
				fmt.Println("c1:", v)
			case v, ok := <-c2:
				if !ok {
					//当 c1或者 c2任一有一个关闭的，那么给发送关闭信号
					o <- true
					fmt.Println("c2  被关闭了！")
					//time.Sleep(1 * time.Second)
					//break 是有效的..
					//break;
				}
				fmt.Println("c2:", v)
				//要是在通个 select 中发送的信号。此处是接收不到的
				//case v,ok := <-o:
				//    if !ok{
				//        break
				//    }
				//
				//    fmt.Println("关闭!",v,ok);
				//    close(o)

			}
		}
	}()

	c1 <- 1
	c2 <- "Hi"
	c1 <- 3
	c2 <- "Hello"

	close(c1)
	//time.Sleep(2 * time.Second)
	close(c2)

	for {
		fmt.Println("kkk")
		oo := <-o
		fmt.Println(oo)
	}

	fmt.Println()
}
