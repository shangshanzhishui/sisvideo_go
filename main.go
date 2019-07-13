package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sis_video_go/api/user"
)
type middleWareHandler struct {
	r *httprouter.Router
}

func MiddleWareHandler(r *httprouter.Router) http.Handler{
	m := middleWareHandler{}
	m.r = r
	return m.r
}

func (m middleWareHandler) ServrHttp(w http.ResponseWriter, r *http.Request){
	m.r.ServeHTTP(w,r)
}

func RegisterHandle() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/login",user.Login())
	return router
}

func main(){
	r := RegisterHandle()
	http.ListenAndServe(":8080",r)
}