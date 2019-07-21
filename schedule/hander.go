package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sis_video_go/api/user"
	"sis_video_go/common"
	"sis_video_go/schedule/sch_db"
	"strconv"
)

func DelVideoRecord(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	username,err := r.Cookie("username")
	id ,err :=user.GetUserId(username.Value)

	vid := p.ByName("video_id")
	aid := p.ByName("authorid")
	athour_id,err:=strconv.Atoi(aid)
	if id == athour_id{
		common.JsonSucess(w,0,200,"")
	}
	if len(vid) == 0{
		common.JsonFail(w,400,"shipin id wei kong")
		return
	}
	err = sch_db.AddVideoDelRecord(vid)
	if err!=nil{
		log.Println(err)
		common.JsonFail(w,500,"shan chu shi bai ")
		return
	}

	return
}
