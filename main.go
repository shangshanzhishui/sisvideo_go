package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sis_video_go/common"
	"sis_video_go/stream_server"

	"sis_video_go/api/user"
)
type middleWareHandler struct {
	r *httprouter.Router
}

type streamMiddleWareHandle struct {
	 r *httprouter.Router
	 l *stream_server.ConnLimiter
}

func NewStreamMiddleWareHandle(r *httprouter.Router ,cc int){
	m := streamMiddleWareHandle{}
	m.r = r
	m.l = stream_server.NewConnLimiter(cc)
}

func (m streamMiddleWareHandle) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if !m.l.GetCon(){
		common.JsonFail(w,http.StatusTooManyRequests,"too many request")
		return
	}
	m.r.ServeHTTP(w,r)
	defer m.l.ReleaseConn()
}

func MiddleWareHandler(r *httprouter.Router) http.Handler{
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	m.r.ServeHTTP(w,r)
}

func RegisterHandle() *httprouter.Router{
	router := httprouter.New()
	//router.POST("/user",CreateUser)
	router.POST("/login",user.Login)
	return router
}

func main(){
	r := RegisterHandle()
	m := MiddleWareHandler(r)
	http.ListenAndServe(":8080",r)
	http.ListenAndServe(":9090",m)

}