package main

import (
    "fmt"
    "time"
)

func main()  {

    //用于数据传输
    c1,c2 := make(chan int),make(chan string)
    //用于信号发送
    o := make(chan bool)

    go func() {
        for {
            select {
                case v,ok := <-c1:
                    if !ok{
                        //当 c1或者 c2任一有一个关闭的，那么给发送关闭信号
                        o <- true
                        time.Sleep(1 * time.Second)
                        break
                    }
                    fmt.Println("c1:",v)
                case v,ok := <-c2:
                    if !ok{
                        //当 c1或者 c2任一有一个关闭的，那么给发送关闭信号
                        o <- true
                        time.Sleep(1 * time.Second)
                        break
                    }
                    fmt.Println("c2:",v)
                case v,ok := <-o:
                    if !ok{
                        break
                    }

                    fmt.Println("关闭!",v,ok);
                    close(o)

            }
        }
    }()

    c1 <- 1
    c2 <- "Hi"
    c1 <- 3
    c2 <- "Hello"

    close(c1)
    close(c2)

    //oo := <- o
    //fmt.Println(oo)

    fmt.Println();
}

