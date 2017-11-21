package main

import (
    "net/http"
    "log"
    "io"
    "strings"
)

func main() {
    http.HandleFunc("/", Cookie)
    http.HandleFunc("/2", Cookie2)

    log.Fatalln(http.ListenAndServe(":8080",nil))
}

func Cookie(w http.ResponseWriter,r *http.Request)  {
    ck := &http.Cookie{
        Name:"myCookie",
        Value:"hello",
        Domain:"localhost",
        Path:"/",
        MaxAge:120,
    }

    http.SetCookie(w,ck)

    ck2,err := r.Cookie("myCookie")

    if err != nil{
        io.WriteString(w,err.Error())
        return
    }

    io.WriteString(w,ck2.Value)
}


func Cookie2(w http.ResponseWriter,r *http.Request)  {
    ck := &http.Cookie{
        Name:"myCookie2",
        Value:"hello world!",
        Domain:"localhost",
        Path:"/",
        MaxAge:120,
    }

    w.Header().Set("Set-Cookie",strings.Replace(ck.String()," ","%20s",-1))

    ck2,err := r.Cookie("myCookie2")

    if err != nil{
        io.WriteString(w,err.Error())
        return
    }

    io.WriteString(w,ck2.Value)
}