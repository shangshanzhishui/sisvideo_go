package video

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"sis_video_go/api/user"
	"sis_video_go/common"
	"sis_video_go/db"
	"sis_video_go/utils"
	"time"
)

type Video_info struct {
	Author_id 	 int 	`json:"author_id"`
	Id  		 string `json:"id"`
	Video_name   string `json:"video_name"`
	Display_time string `json:"display_time"`
}

func AddVideo(aid int,name string)(*Video_info,error){
	id := utils.NewUUID()
	t := time.Now()
	ctime := t.Format("jan 02 2006,15:04:05")
	stmt ,err := db.Db.Prepare(`insert into video_info 
    (id ,author_id,video_name,display_time) values (?,?,?,?)`)
	if err != nil{
		return nil,err
	}
	_, err = stmt.Exec(id,aid,name,ctime)
	if err != nil{
		return nil,err
	}
	res := &Video_info{aid,id,name,ctime}
	defer stmt.Close()
	return res,nil
}
type videoInfo struct {
	Username  string
	Video_name string
}
func AddVIdeoInfo(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	ok,err:=user.VaildUser(w,r)
	if !ok{
		common.JsonSucess(w,"",200,err.Error())
		return
	}
	res,_ := ioutil.ReadAll(r.Body)
	videoBody := &videoInfo{}
	if err := json.Unmarshal(res,videoBody);err !=nil{
		log.Println(err)
		common.JsonFail(w,500,err.Error())
		return
	}
	author_id,err :=user.GetUserId(videoBody.Username)
	if err!=nil{
		common.JsonFail(w,500,err.Error())
		return
	}
	video_info ,err:=AddVideo(author_id,videoBody.Video_name)
	if err!=nil{
		common.JsonFail(w,500,err.Error())
		return
	}else{
		common.JsonSucess(w,video_info,200,"ok")
	}

}

func ListAllVideos(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	ok,err:=user.VaildUser(w,r)
	if !ok{
		common.JsonSucess(w,"",200,err.Error())
		return
	}
	username := p.ByName("username")

}

func GetVideoInfo(id string) (*Video_info,error){
	video_info := Video_info{}
	stmt,err := db.Db.Prepare("select id ,author_id,Video_name,display_time from video_info where id=?")
	if err != nil{
		return nil,err
	}
	err = stmt.QueryRow(id).Scan(&video_info.Id,&video_info.Author_id,&video_info.Video_name,&video_info.Display_time)
	if err != nil && err!=sql.ErrNoRows{
		return nil ,err
	}

	if err == sql.ErrNoRows{
		return nil,nil
	}
	defer stmt.Close()
	return &video_info,nil

}

func DeleteVideoInfo(id string) error{
	stmt,err := db.Db.Prepare("delete from video_info where id=?")
	if err != nil{
		return err
	}
	_,err = stmt.Exec(id)
	if err != nil{
		return err
	}
	defer stmt.Close()
	return nil
}

func GetUserVideos(username string)([]map[string]string, error){
	var video_name  string
	var display_time string
	var m map[string]string
	lv:= []map[string]string{}
	m =make(map[string]string)

	author_id ,err := user.GetUserId(username)
	if err!=nil{
		return nil,err
	}
	stmt,err := db.Db.Prepare("select video_name,dispaly_time from video_info where author_id=?")
	if  err!=nil{
		return nil,err
	}
	rows,err := stmt.Query(author_id)
	if  err!=nil{
		return nil,err
	}

	for rows.Next(){
		err:=rows.Scan(&video_name,&display_time)
		if err != nil{
			return nil,err
		}
		m["video_name"] = video_name
		m["display_time"] = display_time
		lv = append(lv,m)

	}
	return lv,nil
}

func GetAllVideos()([]map[string]string, error){
	var video_name  string
	var display_time string
	var m map[string]string
	lv:= []map[string]string{}
	m =make(map[string]string)



	stmt,err := db.Db.Prepare("select video_name,dispaly_time from video_info")
	if  err!=nil{
		return nil,err
	}
	rows,err := stmt.Query()
	if  err!=nil{
		return nil,err
	}

	for rows.Next(){
		err:=rows.Scan(&video_name,&display_time)
		if err != nil{
			return nil,err
		}
		m["video_name"] = video_name
		m["display_time"] = display_time
		lv = append(lv,m)

	}
	return lv,nil
}
