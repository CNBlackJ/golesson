package main

import "golesson/route"

func main () {
	r := router.Router

	router.SetRouter()

	r.Run(":1234")
}