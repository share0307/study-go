package main

import (
	"os"
	"fmt"
	"net/rpc"
	"log"
)

type Args struct {
	A,B int
}

type Quotation struct {
	Quo,Rem int
}

func main(){
	if len(os.Args) != 2{
		fmt.Println("Usage:",os.Args[0],"server")
		return
	}

	serverAddr := os.Args[1]

	//连接(拨号)到rpc 服务器，使用 Gob 编码的
	client,err := rpc.DialHTTP("tcp",serverAddr+":1234")

	if err!= nil{
		log.Fatalln("dialoing:",err)
		return
	}

	//设置参数
	args := Args{17,8}

	//用于保存结果
	//必须是指针类型
	var reply int
	//简单来说，第一个参数就是在服务端注册得参数得调用路径，第二个参数就是传入得参数，第三个参数是为了保存获取后的数据
	err = client.Call("Math.Multiply",args,&reply)

	if err != nil{
		log.Println("Math error:",err)
	}

	fmt.Println(fmt.Sprintf("Math:%d*%d=%d",args.A,args.B,reply))

	//用于保存结果，必须是指针类型
	var quo Quotation
	err = client.Call("Math.Divide",args,&quo)

	if err != nil{
		log.Println("Math error:",err)
	}

	fmt.Println(fmt.Sprintf("Math:%d/%d=%d",args.A,args.B,quo.Quo))
	fmt.Println(fmt.Sprintf("Math:%d%%%d=%d",args.A,args.B,quo.Rem))

	fmt.Println()


}
