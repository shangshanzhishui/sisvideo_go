package comments

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"sis_video_go/api/user"
	"sis_video_go/common"
)
type co struct {
	Username string `json:"username"`
	Content string `json:"content"`
}
func AddCommentHandler(w http.ResponseWriter, r *http.Request,p httprouter.Params){
	video_id := p.ByName("video_id")
	c := &co{}
	res,_:= ioutil.ReadAll(r.Body)
	err:=json.Unmarshal(res,c)
	author_id ,err:= user.GetUserId(c.Username)
	_,err =AddComment(video_id,author_id,c.Content)
	log.Println(c)
	if err!= nil{
		log.Println(err)
		common.JsonFail(w,500,"评论失败")
	}
	common.JsonSucess(w,"",200,"评论成功")
}

type com struct {
	id string `json:"id"`
}
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request,p httprouter.Params){

	res,_:= ioutil.ReadAll(r.Body)
	co := &com{}

	err:=json.Unmarshal(res,co)

	err=DeleteComment(co.id)
	if err!= nil{
		log.Println(err)
		common.JsonFail(w,500,"删除失败")
	}
	common.JsonSucess(w,"",200,"删除成功")

}