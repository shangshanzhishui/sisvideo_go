package user

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"sis_video_go/common"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	res,_ := ioutil.ReadAll(r.Body)
	user := &User{}
	if err := json.Unmarshal(res,user); err != nil{
		res := common.JsonFail(w,"注册失败")
		fmt.Fprint(w)
		return
	}
	if err := AddUser(user.Username,user.Pwd);err != nil{

	}
}
