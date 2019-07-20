package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sis_video_go/api/user"
	"sis_video_go/api/video"
	"sis_video_go/common"
	"sis_video_go/setting"
	"sis_video_go/utils"
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
	//w.Header().Set("Content-Type","video/rmvb")
	http.ServeContent(w,r,"",time.Now(),video)
	defer video.Close()
}

func UpdateHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){

	//log.Println(string(res))
	username := p.ByName("username")
	log.Println(username)
	athour_id,err:=user.GetUserId(username)
	//if err := r.ParseMultipartForm(setting.UPLOADE_MAX_SIZE);err!= nil{
	//	common.JsonFail(w,http.StatusBadRequest,"上传食品超过1000M")
	//	return
	//}
	image_id := utils.NewUUID()
	video_id := utils.NewUUID()
	image_path := setting.IMAGE_DICTORY + image_id
	video_path := setting.VIDEO_DICTORY + video_id
	//res,err:=ioutil.ReadAll(r.Body)
	file,_ ,err := r.FormFile("file")
	image,_,err := r.FormFile("image")
	video_name:= r.PostForm["name"][0]

	log.Println(video_name)
	if err!= nil{
		log.Println(err)
		//log.Println(string(res))
		log.Println("123")
		common.JsonFail(w,http.StatusInternalServerError,"上传失败")

		return
	}
	video_data,err := ioutil.ReadAll(file)
	image_date,err := ioutil.ReadAll(image)
	if err !=nil{
		log.Printf("读取错误:%v",err)
		common.JsonFail(w,http.StatusInternalServerError,"上传出错")
	}
	_,err= os.Create(image_path)
	if err := ioutil.WriteFile(image_path,image_date,0777) ; err!= nil{
		log.Printf("write file error:%v",err)
		common.JsonFail(w,http.StatusInternalServerError,"shangchuanchuchuo")
		return
	}

	_,err= os.Create(video_path)
	if err := ioutil.WriteFile(video_path,video_data,0777) ; err!= nil{
		log.Printf("write file error:%v",err)
		common.JsonFail(w,http.StatusInternalServerError,"shangchuanchuchuo")
		return
	}
	video.AddVideo(video_id,athour_id,video_name,video_path,image_path)
	w.WriteHeader(http.StatusCreated)
	common.JsonSucess(w,"ok",200,"shangchuanchenggong")
}

func UpdateMuiltHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){

}