package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video/controller"
)

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/v1/user", controller.CreateUser)
	router.POST("/api/v1/user/:username", controller.CreateUserName)
	return router
}

func main() {

	r := RegisterHandles()

	http.ListenAndServe(":8000", r)
}
