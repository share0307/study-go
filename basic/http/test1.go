package main

import (
	"net/http"
	"fmt"
)

type SingleHost struct {
	handler http.Handler
	allowedHost string
}

func (this *SingleHost)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	//fmt.Println("request:",r)
	//fmt.Println("response:",w)
	fmt.Println("Hello World!")
}

func myHandle(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Hello World!"))
}


func main() {
	single := &SingleHost{
		handler: http.HandlerFunc(myHandle),
		allowedHost:"127.0.0.1:8080",
	}

	http.ListenAndServe(":8080",single)
}
