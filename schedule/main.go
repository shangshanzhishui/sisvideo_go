package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sis_video_go/schedule/taskruner"
)

func RegisterHandler() *httprouter.Router{
	router := httprouter.New()
	router.DELETE("/video/:video_id",DelVideoRecord)
	return router
}

func main(){
	go taskruner.WorkStart()
	r := RegisterHandler()
	http.ListenAndServe(":9001",r)
}
