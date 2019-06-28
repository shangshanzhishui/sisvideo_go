package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandle() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/login",Login)
	return router
}

func main(){
	r := RegisterHandle()
	http.ListenAndServe(":8080",r)
}