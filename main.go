package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sis_video_go/api/comments"
	"sis_video_go/api/user"
	"sis_video_go/session"
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

	router.GET("/",user.Register)
	router.POST("/register",user.Register)

	router.POST("/login",user.Login)
	router.GET("/user/:username",user.GetUserInfo)
	router.DELETE("/videos/:video_id/comment",comments.DeleteCommentHandler)
	router.POST("/videos/:video_id/comment",comments.AddCommentHandler)
	router.GET("/videos/:video_id/comments",user.ShowComments)
	router.POST("/videos/:video_id/comments",user.PostComments)
	router.POST("/videos/",user.PostComments)
	return router
}

func main(){
	session.LoadAllSession()
	r := RegisterHandle()
	m := MiddleWareHandler(r)
	http.ListenAndServe(":9001",m)


}