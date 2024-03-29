package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sis_video_go/schedule/taskruner"
)

func RegisterHandler() *httprouter.Router{
	router := httprouter.New()
	router.DELETE("/videos/:video_id/user/:username",DelVideoRecord)
	return router
}

func main(){
	go taskruner.WorkStart()
	r := RegisterHandler()
	http.ListenAndServe("0.0.0.0:9003",r)
}
