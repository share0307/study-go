package main

import "fmt"

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func main() {
	//初始化结构里的匿名结构例子，以下例子正确
	//var a person = person{
	//    Name:"kkk",
	//    Age:20,
	//    Contact: struct{
	//            Phone, City string
	//        }{
	//            Phone: "13178818349",
	//            City: "gz",
	//        }
	//    ,
	//}

	//以下为另一种方法，正确的
	var a *person = &person{
		Name: "kkk",
		Age:  20,
	}

	a.Contact.Phone = "13178818349"
	a.Contact.City = "gz"

	fmt.Println(a)

	fmt.Println("Hello World!")
}
