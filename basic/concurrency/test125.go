package main

import "fmt"

func xrange() chan int{ // xrange用来生成自增的整数
	var ch chan int = make(chan int)

	go func() { // 开出一个goroutine
		//当不填写第二个参数，那么也可以认为是无限循环
		//那么要是信道索要数据，那么循环一次
		//其实说是索要，其实是因为阻塞，因为该 channel 是无缓冲的，然后要是不取数据，那么会一直阻塞死循环，当取完以后，则会进行下一个循环阻塞
		//因为是消费者主动，所以才被认为是阻塞
		for i := 0; ; i++ {
			ch <- i  // 直到信道索要数据，才把i添加进信道
		}
	}()

	return ch
}

func main() {
	//http://blog.csdn.net/skh2015java/article/details/60330975
	//类似生成器..
	generator := xrange()

	for i:=0; i < 10; i++ {  // 我们生成1000个自增的整数！
		fmt.Println(<-generator)
	}
}
