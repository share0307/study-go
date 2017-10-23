package main

import "fmt"

//内嵌接口

//usb 接口，使用 interface 关键字
type USB interface {
	//方法名称  返回类型
	GetName() string
	SetName(name string)

	//嵌入 Connecter 接口
	Connecter
	//必须要完全思想此接口中定义的方法
	//Test() string
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

/*
返回名称方法
*/
func (pc PhoneConnecter) GetName() string {
	return pc.name
}

/*
设置名称方法
*/
func (pc PhoneConnecter) SetName(name string) {
	pc.name = name
}

/*
设置名称方法
*/
func (pc PhoneConnecter) Test() {
	fmt.Println("test!")
}

/*
连接方法
*/
func (pc PhoneConnecter) Connect() {
	fmt.Println(pc.name, " connect ok!")
}

//interface{}为空接口
func Disconnect(usb interface{}) {
	//func Disconnect(usb USB)  {

	var name string
	//usb.(type)只能在 switch 中使用！
	switch usb.(type) {
	case PhoneConnecter:
		name = "PhoneConnecter"
	default:
		fmt.Println("Unknow decive.")
	}

	fmt.Println("Disconnect:", name)
}

func main() {
	//把 a 声明为接口类型，那么 a 只能调用在接口里定义的方法了，结构本身的方法无法调用
	var a USB

	a = PhoneConnecter{name: "sb"}
	a.SetName("kkk")
	a.GetName()
	//把 a 声明为接口类型，那么 a 只能调用在接口里定义的方法了，结构本身的方法无法调用
	//a.Test()

	//此处无法捕抓这个错误 fmt.Println("Where are you,kkk?", usb.(USB))
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()

	Disconnect(a)

	fmt.Println("Hello World!")
}
