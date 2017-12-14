package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//创建一个信号量的chan，缓存为1，（0,1）意义不大
	signalChan := make(chan os.Signal, 1)
	//让进城收集信号量。
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("i am workding!")
	<-signalChan
	ExitFunc()
}
func ExitFunc() {
	fmt.Println("i am exiting!")
}
