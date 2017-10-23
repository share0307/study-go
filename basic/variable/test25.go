package main

import (
	"fmt"
	//"sort"
)

func main() {
	sm := make([]map[int]string, 5)
	for k, v := range sm {
		//这是一个内存拷贝值
		//v = make(map[int]string,1);
		sm[k] = make(map[int]string, 1)
		sm[k][1] = "ok!"
		fmt.Println(k, v)
	}

	fmt.Println(sm)

	fmt.Println("Hello world!")
}
