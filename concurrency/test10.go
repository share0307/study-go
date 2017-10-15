package main

import (
    "fmt"
    "time"
)

func main()  {
    timeout := make (chan bool, 1)
    go func() {
        time.Sleep(2 * time.Second) // sleep one second
        timeout <- true
    }()
    ch := make (chan int)
    //close(ch)

    for {
        select {
        case v,ok := <- ch:
            fmt.Println(v,ok)
        case <- timeout:
            fmt.Println("timeout!")
            close(ch)
            close(timeout)
        }
    }

}

