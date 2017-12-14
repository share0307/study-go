package main

import (
	"fmt"
)

//type Appendmiddleware struct {
//	handler http.Handler
//}

type Sb func(a string,b int)

func (this Sb)Test(a string,b int)  {
	this(a,b)
}

func a(a string,b int)  {
	fmt.Println(a,"--",b)
}

func main(){
	sb := Sb(a)

	sb.Test("Hello",10)

	fmt.Println("Hello World!")
}
