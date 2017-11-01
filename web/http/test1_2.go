package main

import (
    "fmt"
    "net/http"
    "io"
    "log"
)

func init() {
    fmt.Println("a")
}

func init()  {
    fmt.Println("b")
}

func init() {
    fmt.Println("c")
}

func main() {
    mux := http.NewServeMux()

    mux.Handle("/demo",&DemoeHttpMux{})
    mux.Handle("/test",&TestHttpMux{})
    mux.Handle("/",&IndexHttpMux{})

    log.Fatalln(http.ListenAndServe(":8080",mux))
}

type TestHttpMux struct {
    
}

func (this *TestHttpMux)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,"this is TestHttpMux.ServeHTTP!");
}

type DemoeHttpMux struct {

}

func (this *DemoeHttpMux)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,"this is DemoeHttpMux.ServeHTTP!");
    fmt.Println(r.Header)
    fmt.Println(r.Host)
    fmt.Println(r.RequestURI)
    fmt.Println(r.URL)
}

type IndexHttpMux struct {

}

func (this *IndexHttpMux)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,"this is IndexHttpMux.ServeHTTP!");
    fmt.Println(r.Header)
    fmt.Println(r.Host)
    fmt.Println(r.RequestURI)
    fmt.Println(r.URL)
}
