package stream_server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"sis_video_go/common"
	"sis_video_go/setting"
	"time"
)

func StreamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	video_id := p.ByName("video_id")
	Video_dictory := setting.VIDEO_DICTORY + video_id
	video ,err := os.Open(Video_dictory)
	if err != nil{
		common.JsonFail(w,500," error")
	}
	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)
	defer video.Close()
}

func UpdateHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){

}