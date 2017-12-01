package main

import (
	"strings"
	"fmt"
)

func main() {

	var str string = "1,2,3,4,5,6,7,8,9,0,11,12,13,20,30,100"

	/*
		sb := strings.Split(str,",")
		//sb := strings.SplitAfter(str,",")
		//sb := strings.SplitAfterN(str,",",6)
		//sb := strings.SplitN(str,",",6)

		//for k,v := range sb {
		//	fmt.Println("k:",k,"v:",v)
		//}

		sbb := strings.Join(sb,":")

		fmt.Println(sb)
		fmt.Println(sbb)
		*/

		tmp := strings.Replace(str,"1","p",1)


		fmt.Println(str)
		fmt.Println(tmp)


}