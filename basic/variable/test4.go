package main

import "fmt"

//常量
const a int = 100

const b = "test!"

const (
	c = a + 1
	d = c + 1
	e = d - 2
)

const x, y, z int = 1000, 2000, 3000

func main() {
	fmt.Println(a, b, c, d, e, x, y, z)
}
