package main

import (
    "fmt"
    //"sort"
)

func main() {

    D()
    E()
    F()

    fmt.Println("Hello world!");
}


func D(){
    fmt.Println("func D");
}

func E(){
    //使用 recover 是可以做恢复操作的
    defer func() {
        if err := recover();err != nil{
            fmt.Println("Recover in E")
        }
    }();
    //panic 相当于 exit,马上终端执行
    panic("Panic in E")
}

func F(){
    fmt.Println("Func F")
}