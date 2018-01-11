package main

import (
	"fmt"
	"reflect"
)

type Ts struct {
	Username string
}

func main() {
	ts := Ts{
		Username:"kkk",
	}

	test(&ts)

	fmt.Println(ts)
}

func test(ts interface{})  {
	t := reflect.ValueOf(ts)

	f := t.Elem().FieldByName("Username")

	f.SetString("ttt")
}
