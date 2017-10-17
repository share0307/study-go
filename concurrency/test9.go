package main

import (
    "fmt"
    //"time"
    "math/rand"
    "time"
)

func main()  {

    //用于数据传输
    c1,c2 := make(chan int),make(chan string)

    go func() {
        for{
            fmt.Println(" 这里是 c2");
            sb := <- c2
            fmt.Println(sb)
        }
    }()

    go func() {
        for{
            fmt.Println(" 这里是 c1");
            sb := <- c1
            fmt.Println(sb)
        }
    }()

    //使用 select 给 channel 随机发送数据
    for {
        time.Sleep(5000000 * time.Microsecond)
        //fmt.Println("")
        //当 c1以及 c2都关闭掉，那么在 for 循环里的 select 中会一直收到信号，为的是通知一下该频道被关闭了
        select {
            case c1 <- 10:
            case c2 <- "test!":
            case c1 <- test():

        }
    }

    fmt.Println();
}

func test()  int{
    return rand.Int();
}