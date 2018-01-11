package main

import "fmt"

func main() {

	if tmp := B();len(tmp) == 0 {
		fmt.Println(tmp)
	}else if tmp := A();len(tmp) != 0 {
		fmt.Println(tmp)
	}

}

func A() string {
	fmt.Println("A!")
	return "A"
}

func B() string {
	fmt.Println("B!")
	return "B"
}
