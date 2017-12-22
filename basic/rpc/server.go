package main

import (
	"errors"
	"net/rpc"
	"log"
	"net/http"
	"fmt"
)

//此结构体的字段必须得可以 Gob!
type Args struct {
	A,B int
}

//其实是什么类型都没关系
//大概是用于绑定？
type Math int

/**
乘法
1.方法的接受者必须是指针类型的 Math
2......
http://wiki.jikexueyuan.com/project/go-web-programming/08.4.html
 */
func (m *Math)Multiply(args *Args,reply *int) error {
	*reply = args.A * args.B
	fmt.Println("到这里来了:",*reply)
	return nil
}

type Quotient struct {
	Quo,Rem int
}

func (m *Math)Divide(args *Args, quo *Quotient) error {
	if args.A == 0{
		return errors.New("divide by zero!")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func main() {
	var math *Math = new(Math)

	//注册方法
	rpc.Register(math)
	//注册为 http rpc handle,所以下面再次注册 http.ListenAndServe 的时候得 handle 可以设置为 nil
	rpc.HandleHTTP()

	log.Fatalln(http.ListenAndServe(":1234",nil))

}
