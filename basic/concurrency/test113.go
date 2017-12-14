package main

import (
	"fmt"
	"time"
)

//多个接受者
func main() {

	c := make(chan int)

	//这里说明，channel 并不是广播，而是 direct
	go o("sb",c)
	go o("kkk",c)

	for i := 0;i < 10 ;i++  {
		c <- i
	}

	time.Sleep(10 * time.Second)

}

func o(name string,c chan int)  {
	for {
		select {
			case tmp := <- c:
				fmt.Println("name:",name,"--",tmp)
		}
	}
}