package main

import (
    "net/http"
    "log"
    "strings"
    "html/template"
    "fmt"
)

var html string = `
<html>
    <head>
        <title>Hey</title>
    </head>
    <body>
        <form method="post" action="/">
            username:<input name="username" type="text" />
            password:<input name="password" type="password" />
            <button type="submit">提交</button>
        </form>
        <div>
            <p>{{ . }}</p>
        </div>
    </body>
</html>
`

func main()  {
    http.HandleFunc("/",Hey)

    log.Fatalln(http.ListenAndServe(":8080",nil))
}

func Hey(w http.ResponseWriter,r *http.Request)  {
    if strings.ToUpper(r.Method) == "GET"{
        t := template.New("test")
        t.Parse(html)
        t.Execute(w,"Hello World!")
    }else{
        /*
        //获取表单的内容
        r.ParseForm()
        //fmt.Println("form:",form)
        //调用r.Form之前需要先执行一下 r.ParseForm() ，否则 r.Form 是无法获得任何东西的
        //原因是为了为了效率，不会每次都解析一下 form 表单
        fmt.Println("form:",r.Form)
        */

        //使用 FormValue 的话，可以看出，当还没解析的话(也就是说 r.Form 为 nil)，那么会自动调用r.ParseMultipartForm去解析表单，再从中获取返回
        /*
        if r.Form == nil {
            r.ParseMultipartForm(defaultMaxMemory)
        }
        if vs := r.Form[key]; len(vs) > 0 {
            return vs[0]
        }
        return ""
         */
        fmt.Println(r.FormValue("username"))
        fmt.Println(r.FormValue("password"))
    }

    return
}

