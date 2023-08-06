package controller

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "create user controller")
}

func CreateUserName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	io.WriteString(w, username)
}
