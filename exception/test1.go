package main

import "fmt"

func A() (string,bool) {
    fmt.Println("A1");
    defer fmt.Println("A2");

    fmt.Println("B1");
    defer fmt.Println("B2");

    panic("error!!");

    fmt.Println("C1");
    defer fmt.Println("C2");

    return "kkk",true;
}

func recoverErr()  {
    if err:=recover();err!=nil {
        fmt.Println("recoverErr:",err);
    }else{
        fmt.Println("noErr:",err);
    }
}

func main()  {
    defer recoverErr();
    A();

    fmt.Println("Hello World!");
}