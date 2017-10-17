package main

import (
    "fmt"
    "time"
)

func main()  {
    //创建一个用于防止超时的频道
    timeout := make (chan bool, 10)

    //创建一个 ch 的频道，用于数字传输
    ch := make (chan int,10)

    //创建一个 go routine
    go func() {
        //睡眠1秒
        time.Sleep(1 * time.Second) // sleep one second
        //通过 timeout channel 发送 true
        timeout <- true
        ch <- 100
    }()

    go func() {
        for {
            fmt.Println("for 循环:");
            //http://blog.csdn.net/zhonglinzhang/article/details/45913443
            //此处为了验证在 for 循环(无线循环)中，需要至少从 select 中执行一个才能继续下一个循环
            //每次 for 循环只会执行其中一个 case 或者 default，要是需要监听的队列同时发送数据过来，其实也是等执行完一个再在下一次 select 循环中执行
            select {
            case v, ok := <-ch:
                fmt.Println("ch ->");
                if !ok {
                    break;
                }
                fmt.Println(v, ok)
                break;
            case <-timeout:
                fmt.Println("timeout ->");
                time.Sleep(2 * time.Second)
                break;
            //default:
            //    fmt.Println("这里是 default!");
            }

        }
    }()

    //闭包..超时调用模拟定时调用
    var fc func();
    fc = func() {
        time.AfterFunc(1 * time.Second, func() {
            //ch <- 10000000000
            //timeout <- true
            fmt.Println(">>>>>aaaaaaaaaa")
            fc();
        });
    }

    fc();


    //close(timeout)


    time.Sleep(30 * time.Second);


    close(ch)
    close(timeout)

}

