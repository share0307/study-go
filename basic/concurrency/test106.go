package main

import (
	"fmt"
	//"time"
	"time"
)


func main() {
	c := make(chan int)
	//因为一直没有发数据给 flag，所以导致死锁？暂时的说法
	flag := make(chan bool)


	go func() {
		for {
			select {
				case tmp,ok := <- c:
					fmt.Println(tmp)
					if ! ok {
						flag <- true
						//close(flag)
					}
			//case <- time.After(1 * time.Second):
			//	fmt.Println(time.Now())
			}
		}
	}()

	//go func() {
	//	time.Sleep(10 * time.Second)
	//	flag <- false
	//}()

	//c <- 100

	//time.Sleep(5 * time.Second)
	for i := 0;i<101 ;i++  {
		c <- i
	}

	<- flag

	fmt.Println(time.Now())
	fmt.Println()

}