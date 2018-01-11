package main

import "fmt"

func MethodA(name string, f func(a int, b string), sb ...int) {
	fmt.Println("Enter MethodA:", name)
	f(3030, "zdd") // 给f注入参数
	fmt.Println("Exit MethodA:", name)
	fmt.Println(sb[1])
}

func MethodB(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	//把10强制类型转换为 string 类型
	sb := (string)(10)

	fmt.Println(sb)

	d := MethodB
	MethodA("zddhub", d,1,2,3)
}