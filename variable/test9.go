package main

import "fmt"

func main() {
    var(
        a [2]int
        b[2]int
    )
    a = [2]int{1,2};
    b = [2]int{2,2};
    fmt.Println(a);
    fmt.Println(b);

    //类型，数量相同可直接对比
    fmt.Println(a == b);
    fmt.Println(a != [2]int{20,19});


    fmt.Println("Hello world!");
}
