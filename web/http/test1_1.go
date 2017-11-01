package main

import (
    "net/http"
    "io"
    "log"
)

func main()  {
    http.HandleFunc("/hi",SayHi);
    http.HandleFunc("/world",SayWorld);
    handler := http.HandlerFunc(TestHandle)
    http.Handle("/sb",handler)

    log.Fatalln(http.ListenAndServe(":8080",nil))
}

func SayHi(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,r.URL.String())
    io.WriteString(w,"Hi!");
    http.NotFound(w,r)
}

func SayWorld(w http.ResponseWriter,r *http.Request)  {
    log.Println("url not found:",r.RequestURI);
}

func TestHandle(w http.ResponseWriter,r *http.Request)  {
    io.WriteString(w,r.URL.String())
    io.WriteString(w,"TestHandle！");
    //http.NotFound(w,r)
}
