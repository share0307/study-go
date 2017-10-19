package main

import (
    "fmt"
)

func main()  {

    //创建一个channel，并且声明从 channel 中传输的值为 bool 类型
    //第二个参数为缓存的数量，默认为0
    //当 channel 设置了缓存，切并未为堵塞状态下(消息 >= 缓冲值)，程序不管缓冲值中的数据，直接往下直接，也可以直接终止
    c := make(chan bool,1);

    //关键字 go : go 调用方法，就是一个 go routine 了
    go func() {
        fmt.Println("GO GO GO!!!")

        //此处的意思是，从 名为 c 的 channel 中读取值，并且赋值为 sb
        for v := range c{
            fmt.Println("c->",v)
        }
    }()

    //此处的意思是：把 true 这个值通过 名为 c 的 channel 传输给 别的协程
    c <- true
    c <- true
    c <- true

    //关闭 channel
    close(c)

    fmt.Println();
}
