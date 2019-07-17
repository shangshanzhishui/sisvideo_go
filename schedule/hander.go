package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sis_video_go/common"
	"sis_video_go/schedule/sch_db"
)

func DelVideoRecord(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	vid := p.ByName("video_id")
	if len(vid) == 0{
		common.JsonFail(w,400,"shipin id wei kong")
		return
	}
	err := sch_db.AddVideoDelRecord(vid)
	if err!=nil{
		log.Println(err)
		common.JsonFail(w,500,"shan chu shi bai ")
		return
	}
	common.JsonSucess(w,0,200,"")
	return
}
