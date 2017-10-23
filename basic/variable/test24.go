package main

import (
	"fmt"
	//"sort"
)

func main() {

	var m map[int]map[int]string

	//这个地方仅仅是对外层做初始化
	m = make(map[int]map[int]string)
	//对二级元素作初始化
	m[1] = make(map[int]string)

	m[1][10] = "ttt"
	//m[2][100] = "ttt";

	//map，当赋值成功，那么第二个参数(ok)值为 true，失败则为 false
	a, ok := m[2][100]

	if !ok {
		m[2] = make(map[int]string)
		m[2][100] = "ttt"
		a, ok = m[2][100]
	}

	fmt.Println(a, ok)

	fmt.Println("Hello world!")
}
