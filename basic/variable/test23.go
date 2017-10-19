package main

import (
    "fmt"
    //"sort"
)

func main() {

    var m map[int]string = map[int]string{};
    m[1] = "test!";

    //删除 map m 的 key 为1的值，删除后，直接使用则返回为空
    delete(m,1);
    fmt.Println(m[1])
    fmt.Println(m[2])

    fmt.Println("Hello world!");
}
