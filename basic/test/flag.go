package main

import (
	"flag"
	"fmt"
)

var music_file *string = flag.String("file", "musicfile", "Use -file <filesource>")
var test_file *string = flag.String("test", "aa", "Use -test <filesource>")

func main() {
	//flag.String("test","aaa","bbb")
	flag.Parse()
	fmt.Println(*music_file)
	fmt.Println(*test_file)
}