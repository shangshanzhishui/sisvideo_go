package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"sis_video_go/api/user"
	"sis_video_go/common"
	"sis_video_go/db"
	"sis_video_go/schedule/sch_db"
	"strconv"
)

func DelVideoRecord(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	username := p.ByName("username")
	id ,err :=user.GetUserId(username)

	vid := p.ByName("video_id")
	aid,err := GetUserIdByVideoInfo(vid)
	athour_id,err:=strconv.Atoi(aid)
	if id != athour_id{
		common.JsonFail(w,0,"本视频非你所有禁止删除")
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

func GetUserIdByVideoInfo(vid string)(string,error){
	var id string
	stmt,err := db.Db.Prepare("select author_id from video_info where id = ?")
	stmt.QueryRow(vid).Scan(&id)
	if err!=nil{
		return "",err
	}
	return id,nil
}