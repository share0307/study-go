package main

import "fmt"

func main() {
	var a int
	var b float32
	a = 10
	b = 20

	//int + float32是错误的，因为类型不同，go 主张显示转换
	//fmt.Println(a + b);
	fmt.Println(float32(a) + b)

	fmt.Println("Hello World!")
}
