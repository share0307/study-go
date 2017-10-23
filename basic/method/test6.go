package main

import "fmt"

type TT int

func (tt *TT) Increase(num int) {
	//http://blog.csdn.net/sybnfkn040601/article/details/54614798
	fmt.Println(tt)
	//这时候已经改变了原值，因为是引用传递
	//注意：就算TT 是 int 类型的别名，可是因为必须是相同类型才能作对比以及运算，所以还是得把 int 类型强转为 TT 类型
	*tt += TT(num)
	fmt.Println(*tt)
}

func main() {
	t := TT(10)

	t.Increase(100)

	fmt.Println("Hello World!")
}
