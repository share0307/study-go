package main

import (
    "fmt"
    //"sort"
)

func main() {
    //s1 := []int{1,2,3,4,5,6,};
    s1 := make([]int,10,10);

    s2 := []int{7,8,9,10,11,12,13};

    //copy(目标,源),把参数2的数据 copy 到参数1中，按位置替换参数1同位置上的数据
    //当超出参数1的长度时，那么会忽略参数2超长的部分
    //copy(s1,s2);
    //把参数2的指定位置的值copy 到参数1的指定位置
    copy(s1[5:7],s2[2:5]);

    fmt.Println(s2[2:5])
    fmt.Println(s1)

    fmt.Println("Hello world!");
}
