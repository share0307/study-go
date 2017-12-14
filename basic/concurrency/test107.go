package main

import (
	"time"
	"fmt"
)

func main() {
	flag := make(chan bool)


	go func() {
		time.Sleep(5 * time.Second)
		flag <- true
		//close(flag)
	}()

	fg := <- flag

	fmt.Println(fg)
}
