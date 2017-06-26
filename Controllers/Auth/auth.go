package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("Auth | Login")
}
