package main

import (
	"fmt"
	"strconv"
)

func main()  {
	A := A{};
	h := A.Header();
	for i := 0;i<10 ;i++  {
		A.Header()[strconv.Itoa(i)] = "test"
		h[strconv.Itoa(i)] = "test"
		fmt.Println(strconv.Itoa(i))
	}

	fmt.Println(A.Header())
	fmt.Println(h)
}

type A struct {

}

func (this *A)Header() Header {
	return Header{}
}

type Header map[string]string
