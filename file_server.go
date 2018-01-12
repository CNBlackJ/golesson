package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8090", http.FileServer(http.Dir(".")))
}
