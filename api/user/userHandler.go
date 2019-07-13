package user

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"sis_video_go/common"
)

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	res,_ := ioutil.ReadAll(r.Body)
	user := &User{}
	if err := json.Unmarshal(res,user); err != nil{
		common.JsonFail(w,"注册失败")
		return
	}
	if err := AddUser(user.Username,user.Pwd);err != nil{
		common.JsonFail(w,"注册失败")
		return
	}
	data := []string{}
	common.JsonSucess(w,data,"注册成功")
}
