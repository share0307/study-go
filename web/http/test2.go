package main

import (
    "log"
    "net/http"
    "io"
)

func main() {
    mux := http.NewServeMux()

    mux.Handle("/",&MyHandle{})
    mux.Handle("/b",&GoodHandle{})

    mux.HandleFunc("/hello",sayHello)

    err := http.ListenAndServe(":8080",mux);

    if err != nil{
        log.Fatalln(err)
    }


}

type MyHandle struct {

}

func (this *MyHandle)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    io.WriteString(w,"This is the MyHandle's ServerHTTP!uri is "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
    io.WriteString(w,"Hello World! This is the version 2!");
}


type GoodHandle struct {

}

func (this *GoodHandle)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    io.WriteString(w,"This is the GoodHandle's ServerHTTP!uri is "+r.URL.String())
}