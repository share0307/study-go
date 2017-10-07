package main

import "fmt"

//内嵌接口

//usb 接口，使用 interface 关键字
type USB interface {
    //方法名称  返回类型
    GetName()  string
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

type TvConnecter struct {
    name string
}

/*
返回名称方法
 */
func (pc PhoneConnecter)GetName() string  {
    return pc.name;
}

/*
设置名称方法
 */
func (pc PhoneConnecter)SetName(name string)  {
    pc.name = name;
}

/*
返回名称方法
 */
func (tv TvConnecter)GetName() string  {
    return tv.name;
}

/*
设置名称方法
 */
func (tv TvConnecter)SetName(name string)  {
    tv.name = name;
}

/*
设置名称方法
 */
func (pc PhoneConnecter)Test()  {
    fmt.Println("test!");
}

/*
连接方法
 */
func (pc PhoneConnecter)Connect(){
    fmt.Println(pc.name," connect ok!");
}

/*
连接方法
 */
func (tv TvConnecter)Connect(){
    fmt.Println(tv.name," connect ok!");
}

//interface{}为空接口
//func Disconnect(usb interface{})  {
func Disconnect(usb USB)  {
    var name string;
    fmt.Println("Where are you,kkk?", usb.(USB))
    if pc,ok := usb.(PhoneConnecter);ok{
        name = pc.GetName();
        fmt.Println("Disconnected:",name)
        return ;
    }
    fmt.Println("Unknow decive.");
}

func main()  {
    //1.声明一个Phoneconnecter 结构
    tv := TvConnecter{name:"tv"}
    //2.声明一个类型为 Connecter 的接口
    var a USB;

    //3.把 pc 强制转为 Connecter 类型
    //强制转换时是获取到 tv 的内存拷贝，所以当 tv 的值发生改变时，不会影响到 a
    a = USB(tv);
    a.SetName("tv!")

    //强制转换时是获取到 tv 的内存拷贝，所以当 tv 的值发生改变时，不会影响到 a
    tv.name = "aaaaaaaaaaaaaaaaa";

    fmt.Println(a.GetName(),"here!")

    fmt.Println(a);

    Disconnect(a);

    fmt.Println("Hello World!");
}