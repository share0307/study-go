package main

import (
    "fmt"
    //"sort"
)

func main() {
    for i:= 0;i<3;i++{
        v :=1
        fmt.Println(&v);
        fmt.Println(*&v);
    }
    fmt.Println("Hello world!");
}
