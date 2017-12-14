package main

import (
	"time"
	"fmt"
)

func main() {
	//定时调用
	//t := time.NewTicker(1 * time.Second)
	//超时调用
	t := time.NewTimer(1 * time.Second)

	go func() {
		for{
			select {
			case <- t.C:
				fmt.Println(time.Now())
			}
		}
	}()

	time.Sleep(10 * time.Second)
}