package main

import (
    "fmt"
    //"sort"
)

func main() {
    a := [...]int{1,2,3,4,5,6};
    s1 := a[2:5];
    s2 := a[1:3];

    fmt.Println(s1,s2)
    s2 = append(s2,1,2,3,23,53,6,456,567,568,5);
    a[2] = 100;
    fmt.Println(s1,s2)

    fmt.Println("Hello world!");
}
