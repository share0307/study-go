package main

import (
	"reflect"
	"fmt"
)

type Person struct {
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"username"`
}




func main()  {
	person := Person{Name:"kkk"}

	t := reflect.TypeOf(person);

	field := t.Field(0)

	sb := field.Tag.Get("json")

	fmt.Println(sb)


}