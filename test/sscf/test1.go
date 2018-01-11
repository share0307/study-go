package main

import "fmt"

type Business interface {
	handle() func()
}

type A struct {

}

func (this *A)handle()  {
	fmt.Println("this is A!")
}

type B struct {

}

func (this *B)handle()  {
	fmt.Println("this is B!")
}

type C struct {

}

func (this *C)handle()  {
	fmt.Println("this is C!")
}








func ExportBusiness() []Business {
	return []Business{

	}
}






func main() {
	for k,v := range T() {

	}
}

