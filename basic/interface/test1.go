package main

import "fmt"

//usb 接口，使用 interface 关键字
type USB interface {
	//方法名称  返回类型
	Name() string
	Connect()
	//必须要完全思想此接口中定义的方法
	//Test() string
}

type PhoneConnecter struct {
	name string
}

/*
返回名称方法
*/
func (pc PhoneConnecter) Name() string {
	return pc.name
}

/*
连接方法
*/
func (pc PhoneConnecter) Connect() {
	fmt.Println(pc.name, " connect ok!")
}

func Disconnect(usb USB) {
	fmt.Println("Disconnected:", usb.Name())
}

func main() {
	var a USB

	a = PhoneConnecter{}
	a.Name()

	Disconnect(a)

	fmt.Println("Hello World!")
}
