package main

import (
    "fmt"
    //"sort"
)

func main() {


    s1 := make([]int, 3,6);
    //s1 := []int{1,2,3};
    //a := [...]int{100,200,300,400,500,600};
    //s1 := a[2:5];

    fmt.Println(fmt.Sprintf("%v\n%p",s1,s1))

    s2 := append(s1,1,2)

    fmt.Println(s2);
    fmt.Println(len(s2));
    fmt.Println(cap(s2));
    fmt.Println(fmt.Sprintf("%v\n%p",s2,s2))

    fmt.Println("Hello world!");
}
