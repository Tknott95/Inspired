package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	auth "github.com/tknott95/Inspired/Controllers/Auth"
	globals "github.com/tknott95/Inspired/Globals"
)

func main() {

	println("Server running on port: ", globals.PortNumber)
	mux := httprouter.New()
	mux.GET("/", auth.AuthLogin)
	http.ListenAndServe(":"+globals.PortNumber, mux)
}
