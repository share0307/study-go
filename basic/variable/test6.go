package main

import "fmt"

const (
	a = 65
	b = iota
	c = "dd"
	d = iota
	e
)

const (
	x = 1
	y = iota
	z
)

func main() {
	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, z)
}
