package main

import (
    "fmt"
    "time"
)

func main()  {
    //创建一个用于防止超时的频道
    timeout := make (chan bool, 1)
    //创建一个 go routine
    go func() {
        //睡眠1秒
        time.Sleep(1 * time.Second) // sleep one second
        //通过 timeout channel 发送 true
        timeout <- true
    }()

    //创建一个 ch 的频道，用于数字传输
    ch := make (chan int)

    go func() {
        for {
            fmt.Println("Hello World!");
            select {
            case v, ok := <-ch:
                if !ok {
                    fmt.Println("我被关闭拉！");
                    //time.Sleep(5000000 * time.Microsecond)
                    break;
                }
                fmt.Println(v, ok)
                break;
            case <-timeout:
                fmt.Println("timeout!")
                break;
            }
        }
    }()

    //闭包..超时调用模拟定时调用
    var fc func();
    fc = func() {
        time.AfterFunc(5000000 * time.Microsecond, func() {
            timeout <- true

            fc();
        });
    }

    fc();

    //timeout <- true

    //当这个go routine 先关闭了以后，那么善变的 for会一直只执行第一个(ch)，可是关闭 timeout go routine 会报错，因为善变的定时人会推送数据到一个被关闭的 go routine
    close(ch)

    time.Sleep(20 * time.Second);

    close(timeout)

}

