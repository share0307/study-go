package main

import (
    "net/http"
    "io"
    "log"
)

func main()  {
    http.HandleFunc("/hi",SayHi);
    http.HandleFunc("/world",SayWorld);

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
