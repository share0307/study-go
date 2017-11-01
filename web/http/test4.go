package main

import (
    "net/http"
    "io"
    "log"
    "fmt"
)

var mux map[string]func(w http.ResponseWriter,r *http.Request)

func init() {
    mux = make(map[string]func(w http.ResponseWriter,r *http.Request))
    mux["/hello"] = SayHello
    mux["/hi"] = SayHi
    mux["/good"] = SayGood
}

func main()  {
    server := http.Server{

    }

    log.Fatalln(http.ListenAndServe(":8080",mux))
}

type Test struct {

}

func (*Test)ServeHTTP(writer http.ResponseWriter,request *http.Request)  {
    if h,ok := mux[request.URL.String()];ok {
        fmt.Println(request.URL.String())
        h(writer,request)
        return
    }

    io.WriteString(writer,"404 not found!")
}

func SayHello(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,"Hello World!")
}

func SayHi(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,"Hi World!")
}

func SayGood(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,"Good World!")
}


