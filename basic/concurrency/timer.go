package main

import (
	"time"
	"fmt"
)


func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer has expired.")
	}()

	//https://studygolang.com/articles/2915
	//执行timer.Stop()之后，Timer自带的channel并没有关闭，而且这个Timer已经从runtime中删除了
	//timer.Stop()
	timer.Reset(0  * time.Second)
	time.Sleep(60 * time.Second)
}