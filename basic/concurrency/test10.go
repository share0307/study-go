package main

import (
	"fmt"
	//"time"
)

func main() {

	c := make(chan bool)

	//使用空 select 来阻塞 main 方法的结束
	//select {
	//在高版本中的 select{} 已经无效了
	//}

	go func() {
		select {
		case v, ok := <-c:
			fmt.Println(v, ok)
		}
	}()

	close(c)

	fmt.Println()
}
