package main

import "fmt"

func main()  {

	c := make(chan int,10)


	close(c)

	//以下代码证明当channel 为关闭状态，而再次从管道中获取数据，并没有程序通过 channel 传输数据过来，那么强行取的话，那么会取到该类型的默认值
	fmt.Println(<- c)
	fmt.Println(<- c)
	fmt.Println(<- c)

}
