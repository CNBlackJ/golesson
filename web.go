package main

import (
	"github.com/julienschmidt/httprouter"
	"golesson/controller"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/hello", helloController.HelloHandler)
	router.GET("/hello/:name", helloController.EchoHandler)

	log.Fatal(http.ListenAndServe(":8090", router))
}
