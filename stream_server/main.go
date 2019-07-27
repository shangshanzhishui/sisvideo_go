package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sis_video_go/common"
)

type streamMiddleWareHandle struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewStreamMiddleWareHandle(r *httprouter.Router ,cc int) http.Handler{
	m := streamMiddleWareHandle{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}


func (m streamMiddleWareHandle) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if !m.l.GetCon(){
		common.JsonFail(w,http.StatusTooManyRequests,"too many request")
		return
	}
	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}

func registerHandler() *httprouter.Router{
	router := httprouter.New()
	router.GET("/video/:video_id",StreamHandler)
	router.POST("/video/:username",UpdateHandler)
	return router
}

func main(){
	r:= registerHandler()
	//sm:= NewStreamMiddleWareHandle(r,2)
	http.ListenAndServe("0.0.0.0:9002",r)
}