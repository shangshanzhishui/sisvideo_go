package user

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"sis_video_go/common"
	"sis_video_go/session"
)
type user2 struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}
func Register(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	res,_ := ioutil.ReadAll(r.Body)
	fmt.Println(string(res))

	user := &user2{}
	if err := json.Unmarshal(res,user); err != nil{

		common.JsonFail(w,500,"注册失败")
		fmt.Println(user)
		return
	}
	fmt.Println(12)
	fmt.Println(user)


	if err := AddUser(user.Username,user.Pwd);err != nil{
		common.JsonFail(w,500,"注册失败")
		return
	}
	author_id,err := GetUserId(user.Username)
	if err != nil{
		common.JsonFail(w,500,"注册失败")
		return
	}
	id := session.CreateSessionId(author_id)
	data := []map[string]string{{"id":id}}
	cookie := http.Cookie{Name:"session",Value:id}
	http.SetCookie(w,&cookie)
	common.JsonSucess(w,data,200,"注册成功")

	return
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	res,_ := ioutil.ReadAll(r.Body)

	log.Println(res)
	ubody := &user2{}
	if err := json.Unmarshal(res,ubody);err != nil{
		log.Println(err)
		common.JsonFail(w,500,"denglu shibai")
		return
	}
	username := p.ByName("username")
	pwd,err := GetUser(username)
	if err !=nil || len(pwd) ==0 || pwd != ubody.Pwd{
		log.Println(err)
		common.JsonFail(w,500,"denglu shibai")
		return
	}
	session_id,err := session.CreateSessionIdByUname(ubody.Username)
	if err!=nil{
		log.Println(err)
		common.JsonFail(w,500,"denglu shibai")
		return
	}
	common.JsonSucess(w,session_id,200,"ok")
}

func GetUserInfo(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	ok,err:=VaildUser(w,r)
	if !ok{
		common.JsonFail(w,200,err.Error())
	}
	username := p.ByName("username")
	author_id ,err:= GetUserId(username)
	if err!= nil{
		common.JsonFail(w,200,"username is error")
	}
	common.JsonSucess(w,author_id,200,"ok")


}


func ListAllVideos(w http.ResponseWriter,r *http.Request,p httprouter.Params){

}


func ShowComments(w http.ResponseWriter,r *http.Request,p httprouter.Params){

}


func PostComments(w http.ResponseWriter,r *http.Request,p httprouter.Params){

}

