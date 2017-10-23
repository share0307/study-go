package main

import "fmt"

func main() {

	//var a [2]int;
	//var b [2]int;

	//完整写法
	//var a [2]int= [2]int{1};
	//a := [2]int{1:1,0:2};
	//...为省略数组长度
	a := [...]int{20: 1, 0: 2}

	//元素个数以及类型必须得相等才能赋值
	//a = b;

	fmt.Println(a)

	fmt.Println("hello world!")
}
