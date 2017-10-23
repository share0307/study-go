package main

import (
	"fmt"
	//https://studygolang.com/articles/8742
	"reflect"
)

func main() {

	x := 10
	x_point := &x

	fmt.Println("x_old_point,", x_point)
	v := reflect.ValueOf(&x)
	//
	//当ValueOf中传入的是指针时，通过 Elem() 方法取得指针的值，ValueOf 传入的是值则无法使用，那就是报错了
	//Elem():如果是指针，则获取其所指向的元素
	//v.Elem().SetInt(321)
	v = v.Elem()

	fmt.Println("v:", &v)

	//v := reflect.ValueOf(x)
	////无法直接设置值
	//v.SetInt(321)

	//fmt.Println("type:", v.Type())

	fmt.Println(x)

	//注意，断言并不是对直接对任一某个变量进行断言，而是对 接口类型的变量 进行断言
	//(func(x interface{}) {
	//    x_type, ok := x.(int)
	//
	//    if ok {
	//        panic("error!")
	//    }
	//    fmt.Println(x, x_type)
	//})(x);

	fmt.Println("Hello World!")
}
