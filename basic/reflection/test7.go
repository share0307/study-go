package main

import (
	"fmt"
	//https://studygolang.com/articles/8742
	//"reflect"
	//"reflect"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

/*
say hello!
*/
func (this User) Hello(name string) {
	fmt.Println("Hello ", name, ", my name is ", this.Name)
}

func main() {
	u := User{1, "kkk", 100}
	//u.Hello("sb");

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	//参数，用于传递给需要调用的方法，然而 传进去的参数都需要使用 reflect.ValueOf 包装一下，因为参数就是一个 类型为 reflect.Value 的 切片 slice
	args := []reflect.Value{reflect.ValueOf("sb")}

	//调用方法
	mv.Call(args)

	fmt.Println("Hello World!")
}
