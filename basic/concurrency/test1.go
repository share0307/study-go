package main

import (
	"fmt"
	"time"
)

func main() {

	//关键字 go : go 调用方法，就是一个 go routine 了
	go Go()

	//睡眠 10 秒
	//查一下 Sleep() 的参数
	time.Sleep(2 * time.Second)

	fmt.Println("Hello World!~")
}

/*
go routine 输出一句话所调用的方法
*/
func Go() {
	fmt.Println("GO GO GO!!!")
}
