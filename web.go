package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Hello, world!\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "Hello, "+params.ByName("name")+"\n")
}

func main() {
	router := httprouter.New()
	router.GET("/hello", helloHandler)
	router.GET("/hello/:name", echoHandler)

	log.Fatal(http.ListenAndServe(":8090", router))
}
