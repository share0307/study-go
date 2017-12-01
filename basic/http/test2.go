package main

import (
	"net/http"
	"log"
)


type SingleHost struct {
	//1.类型为接口，为的是让此结构可通过 handler 此字段去使用该接口类的方法
	//2.灵活，可以把通用实现了此结构的结构体
	handler http.Handler
	//允许的域名
	host string
	//ServeHTTP func()
}

func (this *SingleHost)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	this.handler.ServeHTTP(w,r)
}

func myHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello World!"))
}

func main()  {
	single := &SingleHost{
		//使用 http.HandlerFunc 把普通方法 myhandle 强转为 http.HandlerFunc 类型
		handler:http.HandlerFunc(myHandle),
		host:"127.0.0.1:8080",
	}

	log.Fatalln(http.ListenAndServe(":8080",single))
}
