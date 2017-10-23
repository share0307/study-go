package main

import "fmt"

//当为匿名字段时(不给字段命名)，在声明时需要严格按照顺序赋值(只要其中有一个是匿名字段都需要按照此规则)
type person struct {
	Name string
	Age  int
}

type person2 struct {
	Name string
	Age  int
}

func main() {

	//以下为另一种方法，正确的
	var a person = person{"kkk", 20}
	//var b person2 = person2{"kkk", 20,}
	var b person = person{"kkk", 20}
	var c person = person{"kkk", 30}

	//相同的结构体可以进行 等于或者不等于 对比，不同的则直接报错(编译都通不过)
	fmt.Println(a == b)

	//当相同的结构体但是数据不同，也并不相等
	fmt.Println(a == c)

	fmt.Println("Hello World!")
}
