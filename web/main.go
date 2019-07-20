package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router{
	router := httprouter.New()

	router.GET("/",indexHandler)
	router.GET("/register",RegiserHandler)
	router.GET("/login",loginHandler)
	router.GET("/upload/:username",uploadHandler)

	router.POST("/register",proxyServeHandler)
	router.POST("/login",proxyServeHandler)
	router.POST("/video/:username",proxyStreamHandler)

	//代理stream——server
	router.POST("/videos",proxyStreamHandler)
	//
	router.GET("/user",userHandler)
	router.ServeFiles("/statics/*filepath",http.Dir("./templates"))

	return router

}
func main(){
	r := RegisterHandler()
	http.ListenAndServe("10.13.165.33:9000",r)
}

