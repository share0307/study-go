package main

import "fmt"

var ch chan int = make(chan int)

func foo(id int) {
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Println(<- ch)
	}
}