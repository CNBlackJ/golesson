package helloController

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Hello, world!\n")
}

func EchoHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "Hello, "+params.ByName("name")+"\n")
}
