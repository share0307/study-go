package main

import (
	"sync"
	"fmt"
)

func main()  {
	//http://www.baiyuxiong.com/?p=913
	//http://blog.haohtml.com/archives/15583
	//WaitGroup在go语言中，用于线程同步，单从字面意思理解，wait等待的意思，group组、团队的意思，WaitGroup就是指等待一组，等待一个系列执行完成后才会继续向下执行。
	//创建同步任务组
	wg := sync.WaitGroup{}
	//设置此组任务为10个一组
	wg.Add(10)

	for i := 0; i<12;i++  {
		go func(wg *sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(&wg)
	}
	//wg.Done()
	//等待所有任务结束，才再次往下执行
	wg.Wait()


}



















//func main() {
//	c := make(chan int,1)
//
//	//for i:= 0;i<2 ;i++  {
//	//	c <- i
//	//}
//
//	c <- 1
//	//c <- 2
//
//	close(c)
//
//}
