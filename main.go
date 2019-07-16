package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"

	"sis_video_go/api/user"
)
type middleWareHandler struct {
	r *httprouter.Router
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
	v,err := os.Open("./videos/123.rmvb")
	if err != nil{
		fmt.Print(err)
	}
	fmt.Println(v)
	r := RegisterHandle()
	m := MiddleWareHandler(r)
	http.ListenAndServe(":8080",m)


}