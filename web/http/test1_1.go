package main

import (
	"net/http"
	"io"
	"log"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!");
	});
	log.Fatal(http.ListenAndServe(":8080", nil))

}