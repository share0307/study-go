package main

import (
	"time"
	"fmt"
)

func main() {
	//以下这段代码要说的是，当队列中无缓冲(缓冲值为0)或者说缓冲满了，再次往 channel 中丢数据的话，会报错
	a := 5
	b := 7
	c := make(chan int,a)

	go func() {
		tmp := <- c

		fmt.Println(tmp)
	}()

	for i:=0;i<b ;i++  {
		c <- i
	}


	time.Sleep(10 * time.Second)

}
