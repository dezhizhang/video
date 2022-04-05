package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"<h1>hello world")
}

func main() {
	http.HandleFunc("/",home)
	http.ListenAndServe(":8000",nil)
}
