package main

import (
	"fmt"
	"net/http"

	auth "github.com/tknott95/Inspired/Controllers/Auth"
	globals "github.com/tknott95/Inspired/Globals"
)

func main() {
	http.HandleFunc("/", auth.AuthLogin)
	fmt.Println("Hello, Trevsauce!")
	http.ListenAndServe(":"+globals.PortNumber, nil)
}
