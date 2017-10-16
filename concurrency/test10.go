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
        time.AfterFunc(1 * time.Second, func() {
            timeout <- true

            fc();
        });
    }

    fc();


    //close(timeout)


    time.Sleep(10 * time.Second);


    close(ch)
    close(timeout)

}

