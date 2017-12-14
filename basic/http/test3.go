package main

import "net/http"

type AppendMiddle struct {
	handler http.Handler
}

func (this *AppendMiddle)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	this.handler.ServeHTTP(w,r)
	w.Write([]byte("hey here!"))
}

func myHandle(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Hello World!"))
}

func main()  {
	mid := &AppendMiddle{http.HandlerFunc(myHandle)}
	http.ListenAndServe(":8080",mid)
}