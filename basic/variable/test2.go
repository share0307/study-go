package main

import "fmt"

import (
	"math"
)

type (
	//别名 类型
	test_uint8 uint8
	文本         string
)

func main() {
	//var a [10]bool;
	fmt.Println(math.MaxInt8)

	//var b 文本 = "啊啊啊啊";
	//fmt.Println(b);

	//b：变量不得重复声明
	//var a,b,c,d = 0,1,2,3;
	//var a,b,c,d int= 0,1,2,3;
	a, b, _, d := 0, 1, 2, 3

	fmt.Println(a, b, d)

	var floata float32 = 100.1

	fmt.Println(floata)

	intb := int(floata)

	fmt.Println(intb)
}
