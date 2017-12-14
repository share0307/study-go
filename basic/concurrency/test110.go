package main

import (
	"time"
	"fmt"
)

func main() {

	b := make(chan string,1000)
	c := make(chan int,1000)

	go func() {
		for{
			select {
				case tmp := <- c:
					time.Sleep(1 * time.Second)
					fmt.Println("int:",tmp)
				case tmp := <- c:
					time.Sleep(1 * time.Second)
					fmt.Println("string:",tmp)
				case t := <- time.After(1 * time.Microsecond):
					fmt.Println("t:",t)
				
			}
		}
	}()


	go func() {
		for i:=0;i<10000 ;i++  {
			c <- i
			//fmt.Println("i:",i)
		}
	}()

	go func() {
		for j:=10000;j<100000 ;j++  {
			c <- j
			//fmt.Println("j:",j)
		}
	}()

	go func() {
		for j:=10000;j<100000 ;j++  {
			b <- string(j)
			//fmt.Println("j:",j)
		}
	}()

	time.Sleep(time.Second * 10)

}