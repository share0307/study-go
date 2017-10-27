package main

import (
    "net/http"
    "fmt"
    "log"
)

func sayHello(w http.ResponseWriter,r *http.Request)  {

    fmt.Fprint(w,"Hello World!!");

    fmt.Println("Hello world!");
}

func main()  {
    http.HandleFunc("/",sayHello)
    http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer,"我勒个去!test!");
        fmt.Fprint(writer,"Current Url Is:",request.URL.Host,request.URL.Port(),request.RequestURI)
        request.ParseForm()
        for k,v := range request.Form{
            fmt.Fprint(writer,"key ",k,":",v)
        }
    })

    err := http.ListenAndServe(":8080",nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}