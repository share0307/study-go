package main

import (
	"fmt"
	"sync"
)

func main() {

	//http://www.baiyuxiong.com/?p=913
	//http://blog.haohtml.com/archives/15583
	//WaitGroup在go语言中，用于线程同步，单从字面意思理解，wait等待的意思，group组、团队的意思，WaitGroup就是指等待一组，等待一个系列执行完成后才会继续向下执行。
	wg := sync.WaitGroup{}
	//添加或者减少等待goroutine的数量
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	//Wait:执行阻塞，直到所有的WaitGroup数量变成0才会继续往下执行代码
	wg.Wait()

	fmt.Println("Hello World!")
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1

	for i := 0; i < 100000; i++ {
		a += i
	}

	fmt.Println(index, a)

	//Done:相当于Add(-1)
	wg.Done()
}
