package main

import (
	"fmt"
	//"time"
	"time"
)

func main() {

	c := make(chan bool)

	//使用空 select 来阻塞 main 方法的结束
	//select {
	//在高版本中的 select{} 已经无效了
	//}

	//go func() {
	select {
	case v, ok := <-c:
		fmt.Println(v, ok)
		break
	case <-time.After(10 * time.Second):
		//After(d Duration) <-chan Time
		//看一下扇面的注释，因为 time.After 实际上也是返回一个 channel
		fmt.Println("10秒，时间到了!")
	}
	//}()

	close(c)

	fmt.Println()
}
