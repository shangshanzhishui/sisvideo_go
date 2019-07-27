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
	router.GET("/video/:video_id",proxyStreamHandler)
	router.GET("/videos/:video_id",videoHandler)
	router.GET("/user/:username",MyHomeHandler)

	//router.GET("/videos/video_id/comments",videoCommentHandler)


	router.POST("/register",proxyServeHandler)
	router.POST("/login",proxyServeHandler)
	router.POST("/video/:username",proxyStreamHandler)
	router.POST("/videos/:video_id/comment",proxyServeHandler)
	router.DELETE("/videos/:video_id/comment",proxyServeHandler)
	router.DELETE("/videos/:video_id/user/:username",proxyScheHandler)

	//代理stream——server
	router.POST("/videos",proxyStreamHandler)
	//
	router.GET("/user",userHandler)
	router.ServeFiles("/statics/*filepath",http.Dir("./templates"))

	return router

}
func main(){
	r := RegisterHandler()
	http.ListenAndServe("0.0.0.0:9000",r)
}

