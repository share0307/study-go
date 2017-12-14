package main

import (
	//"fmt"
	//"time"
)

//import "time"

func main() {

	c := make(chan bool,1)


	//go func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("Hello World!!")
	//}()

	c <- true
	c <- true

	<- c



}