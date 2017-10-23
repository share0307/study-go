package main

import "fmt"

//类型不能不得直接使用 原类型的方法
type TZ int

func (tz *TZ) test() {
	fmt.Println(tz)
}

func main() {
	i := int(10)

	tz := TZ(10)
	tz.test()

	fmt.Println("Hello World!")
}
