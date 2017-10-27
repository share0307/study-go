package main

import (
    "io"
    "log"
    "net/http"
    "fmt"
)

func main() {
    //设置路由
    //http.HandleFunc(路由,执行方法)
    http.HandleFunc("/",SayHello);

    //第一个参数是监听的地址，不写 ip 默认就是本机 ip
    //第二个参数是处理的句柄方法，要是使用默认的话，而传入 nil
    //返回值是一个错误类型的值，当没错误则为 nil
    err := http.ListenAndServe(":8080",nil);

    if err != nil{
        log.Fatalln(err);
    }

    fmt.Println("ok 了哦！")
}

/**
1.从资源角度理解，http.Request的参数都是一致从 http 包底层中获取的，所以要是按值传递每次都会把 http.Request 类型都拷贝传递，是一种资源浪费
2.而 http.ResponseWriter 是一个接口，*http.Request是一个类型的接口
 */
func SayHello(w http.ResponseWriter, r *http.Request)  {
    //简单来说就是通过 w http.ResponseWriter 把 字符串写入！
    io.WriteString(w,"Hello World! This is the version 1!");
}