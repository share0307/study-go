package main

import (
	"net/http"
	"net/http/httptest"
	"fmt"
	//"reflect"
)

type ModifyMiddle struct {
	handler http.Handler
}

func (this *ModifyMiddle)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	//获取一个 rec 对象
	rec := httptest.NewRecorder()

	//把 rec 对象传入到返回总,然后 因为返回的是一个 rec 引用类型，所以可以把赋予输出之给拿到去
	this.handler.ServeHTTP(rec,r)


	var str []string = []string{"aaaaa12345"}
	for k,v := range rec.Header(){
		fmt.Println(k,"--",v)
		//因为并没有真正输出到 w(http.ResponseWriter),所以需要循环输出
		w.Header()[k] = v
		//fmt.Println(reflect.TypeOf(v))
		fmt.Println(w.Header())
	}

	w.Header()["test"] = str

	w.Header().Set("test","kkk")
	w.WriteHeader(418)
	w.Write(rec.Body.Bytes())
	w.Write([]byte("Hello World!!"))
}

func myHandle(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Hello KKK!"))
}

func main()  {
	mid := &ModifyMiddle{http.HandlerFunc(myHandle)}
	http.ListenAndServe(":8080",mid)
	fmt.Println("Hello World!")
}
