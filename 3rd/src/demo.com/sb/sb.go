package sb

import "fmt"

func init() {
	fmt.Println("Hello World!")

	for i:=0;i<10 ;i++  {

		//循环里创建是可行的
		var t int
		fmt.Println(t)
	}

	//for i:=0;i<10 ;i++  {
	//
	//	if i>0{
	//		//不可见了
	//		fmt.Println(b)
	//	}
	//	//循环里创建是可行的
	//	var b int
	//	fmt.Println(b)
	//}

}
