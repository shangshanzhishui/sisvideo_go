package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sis_video_go/common"
	"sis_video_go/setting"
	"time"
)

func StreamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	video_id := p.ByName("video_id")
	Video_dictory := setting.VIDEO_DICTORY + video_id
	//utils.GetAllFil("./")
	fmt.Println(Video_dictory)
	//fmt.Println(video_id)
	video ,err := os.Open(Video_dictory)
	log.Print(video)
	if err != nil{
		common.JsonFail(w,500,"文件出错")
		log.Println(err)
	}
	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)
	defer video.Close()
}

func UpdateHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	if err := r.ParseMultipartForm(setting.UPLOADE_MAX_SIZE);err!= nil{
		common.JsonFail(w,http.StatusBadRequest,"上传食品超过100M")
		return
	}

	file,_ ,err := r.FormFile("file")
	if err!= nil{
		common.JsonFail(w,http.StatusInternalServerError,"上传失败")
		return
	}
	data,err := ioutil.ReadAll(file)
	if err !=nil{
		log.Printf("读取错误:%v",err)
		common.JsonFail(w,http.StatusInternalServerError,"上传出错")
	}

	f := p.ByName("video_id")
	if err := ioutil.WriteFile(setting.VIDEO_DICTORY+f,data,0777) ; err!= nil{
		log.Printf("write file error:%v",err)
		common.JsonFail(w,http.StatusInternalServerError,"shangchuanchuchuo")
		return
	}
	w.WriteHeader(http.StatusCreated)
	common.JsonSucess(w,"ok",200,"shangchuanchenggong")
}