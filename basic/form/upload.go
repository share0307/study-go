package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/test",Tpl)

	http.HandleFunc("/upload",Upload)

	log.Fatal(http.ListenAndServe(":8089",nil))

	fmt.Println("Hello World!")
}

func Tpl(w http.ResponseWriter,r *http.Request)  {
	var tpl string = `
<html>
    <head>
        <title>Hey</title>
    </head>
    <body>
        <form method="post" action="/upload" enctype="multipart/form-data">
            username:<input name="username" type="text" />
            file:<input name="file" type="file" />
            <button type="submit">提交</button>
        </form>
        <div>
            <p>{{ . }}</p>
        </div>
    </body>
</html>
`

t := template.New("test")
t.Parse(tpl)
t.Execute(w,nil)

}

func Upload(w http.ResponseWriter,r *http.Request)  {
	file,reader,_ := r.FormFile("file")

	fmt.Println(file)
	fmt.Println(reader)


	path,_ := filepath.Abs(reader.Filename)
	fmt.Println("path:",path)
	os.Create(path)
	os.Remove(path)
}
