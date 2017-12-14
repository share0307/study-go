package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int,5)

	go l(c)

	//time.Sleep(1 * time.Second)
	for i:=0;i<10 ;i++  {
		c <- i
	}

}

func l(c chan int)  {
	for i:=0;i<10 ;i++  {
		time.Sleep(1 * time.Second)
		fmt.Println(<- c)
	}
}
