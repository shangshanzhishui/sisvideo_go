package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router{
	router := httprouter.New()

	router.GET("/",homeHandler)

	//router.POST("/user",userHandler)
	//
	//router.POST("/user",userHomeHandler)
	//
	router.GET("/user",userHandler)
	router.ServeFiles("/statics/*filepath",http.Dir("./templates"))
	return router

}
func main(){
	r := RegisterHandler()
	http.ListenAndServe(":9002",r)
}

