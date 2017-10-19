package main

import "fmt"

func main() {
    a := [10]int{};
    a[1] = 2;
    fmt.Println(a)

    b := &[10]int{};
    b[1] = 30;
    fmt.Println(b);

    p := new([10]int);
    p[1] = 30;
    fmt.Println(p)

    //b 跟 p 效果是一致的吗？

    fmt.Println(b == p);        //false,原因是引用的地址不同


    fmt.Println("Hello world!");
}
