package main

import (
    "fmt"
)

func main()  {

    //创建一个channel，并且声明从 channel 中传输的值为 bool 类型
    c := make(chan bool);

    //关键字 go : go 调用方法，就是一个 go routine 了
    //c chan bool是按引用传参
    go func(c chan bool) {
        fmt.Println("GO GO GO!!!")

        //此处的意思是：把 true 这个值通过 名为 c 的 channel 传输给 主线程
        c <- true

        //最终结果必须得关闭 channel，否则会导致死锁的问题
        close(c)
    }(c)

    for v := range c {
        fmt.Println(v)
    }

    //以下代码证明当channel 为关闭状态，而再次从管道中获取数据，并没有程序通过 channel 传输数据过来，那么强行取的话，那么会取到该类型的默认值
    /*
    var sb bool
    sb = <- c
    fmt.Println("sb1:",sb)
    sb = <- c
    fmt.Println("sb2:",sb)
    */

    fmt.Println();
}
