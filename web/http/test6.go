package main

import (
    "fmt"
    "os"
    "net/http"
    "log"
)



func main() {
    wd,ok := os.Getwd()

    if ok != nil {
        fmt.Println("err!")
    }

    fmt.Println("当前目录为:",wd)

    //htp.StripPrefix返回的是一个 httphandle
    static_handle := http.StripPrefix("/static/",http.FileServer(http.Dir(wd)))
    fmt.Println(http.FileServer(http.Dir(wd)))

    http.Handle("/static/",static_handle)

    log.Fatalln(http.ListenAndServe(":8080",nil));


}