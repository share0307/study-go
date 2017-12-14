package main

import (
	"fmt"
	//"time"
	"time"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	c := make(chan int)
	//创建一个系统信号监听得渠道
	sg := make(chan os.Signal)
	//设置监听的型号
	signal.Notify(sg,syscall.SIGINT,syscall.SIGALRM)

	tag:

	go func() {
		for {
			select {
				case tmp,ok := <- c:
					fmt.Println(tmp)
					if ! ok {
						//os.Exit(0)
						//close(flag)
					}
			//case <- time.After(1 * time.Second):
			//	fmt.Println(time.Now())
			}
		}
	}()

	//go func() {
	//	time.Sleep(10 * time.Second)
	//	flag <- false
	//}()

	//c <- 100

	//time.Sleep(5 * time.Second)
	for i := 0;i<101 ;i++  {
		c <- i
	}
	//close(c)

	<- sg

	fmt.Println(time.Now())
	goto tag
	fmt.Println()
}