package main

import "fmt"

func main() {
    //多位数组
    //非顶级的数组元素不能使用 ...
    a :=[...][3]int{
        {1,2,3},
        {2,},        //最后的元素要不以 } 结尾，要不以,结尾
    }

    fmt.Println(a);

    fmt.Println("Hello world!");
}
