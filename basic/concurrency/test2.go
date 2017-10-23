package main

import (
	"fmt"
)

func main() {

	//创建一个channel，并且声明从 channel 中传输的值为 bool 类型
	c := make(chan bool)

	//关键字 go : go 调用方法，就是一个 go routine 了
	go func() {
		fmt.Println("GO GO GO!!!")

		//此处的意思是：把 true 这个值通过 名为 c 的 channel 传输给 主线程
		c <- true
	}()

	//此处的意思是，从 名为 c 的 channel 中读取值，并且赋值为 sb
	var sb bool = <-c

	//关闭 channel
	close(c)

	fmt.Println(sb)

	fmt.Println()
}
