package main

import (
	"fmt"
	"time"
	"sync"
)

var a int

func main() {
	ch := make(chan int)
	go foo()
	go foo()
	go foo()
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 10
	}()
	<-ch
}

func foo() {
	mx := new(sync.Mutex)
	mx.Lock()
	defer mx.Unlock()
	a++
	fmt.Println(a)
}
