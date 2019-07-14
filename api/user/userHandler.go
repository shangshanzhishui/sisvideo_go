package user

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"sis_video_go/common"
	"sis_video_go/session"
)
type user2 struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}
func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	res,_ := ioutil.ReadAll(r.Body)
	user := &user2{}
	if err := json.Unmarshal(res,user); err != nil{
		common.JsonFail(w,"注册失败")
		return
	}
	if err := AddUser(user.Username,user.Pwd);err != nil{
		common.JsonFail(w,"注册失败")
		return
	}
	author_id,err := GetUserId(user.Username)
	if err != nil{
		common.JsonFail(w,"注册失败")
		return
	}
	id := session.CreateSessionId(author_id)
	data := []map[string]string{{"id":id}}
	common.JsonSucess(w,data,"注册成功")
	return
}
