package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!!!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8090", nil)
}
