package main

import "fmt"

func main() {

	c := make(chan int,2)
	qa := make(chan bool,1)

	go func(c chan int) {
		fmt.Println("go go go!")
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		c <- 10
		//qa <- true
		close(qa)
		close(c)
	}(c)

	//go func(c chan int) {
	//	fmt.Println("go go go!")
	//	c <- 20
	//	c <- 20
	//	c <- 20
	//	c <- 20
	//	c <- 20
	//	c <- 20
	//	c <- 20
	//	c <- 20
	//	c <- 20
	//	//close(c)
	//}(c)

	go func() {
		for v := range c {
			fmt.Println("value:",v)
		}
	}()

	//go func() {
	//	for v := range qa {
	//		fmt.Println("value:",v)
	//	}
	//}()

	<- qa

	//close(qa)
}
