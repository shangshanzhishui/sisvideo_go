package video

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"sis_video_go/api/comments"
	"sis_video_go/api/user"
	"sis_video_go/common"
	"sis_video_go/db"
	"time"
)

type Video_info struct {
	Author_id 	 int 	`json:"author_id"`
	Id  		 string `json:"id"`
	Video_name   string `json:"video_name"`
	Display_time string `json:"display_time"`
}

func AddVideo(id string,aid int,name string,video_path string,image_path string)(*Video_info,error){

	t := time.Now()
	ctime := t.Format("jan 02 2006,15:04:05")
	stmt ,err := db.Db.Prepare(`insert into video_info 
    (id ,author_id,video_name,display_time,video_path,image_path) values (?,?,?,?,?,?)`)
	if err != nil{
		return nil,err
	}
	_, err = stmt.Exec(id,aid,name,ctime,video_path,image_path)
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

//func AddVIdeoInfo(w http.ResponseWriter,r *http.Request,p httprouter.Params){
//	ok,err:=user.VaildUser(w,r)
//	if !ok{
//		common.JsonSucess(w,"",200,err.Error())
//		return
//	}
//	res,_ := ioutil.ReadAll(r.Body)
//	videoBody := &videoInfo{}
//	if err := json.Unmarshal(res,videoBody);err !=nil{
//		log.Println(err)
//		common.JsonFail(w,500,err.Error())
//		return
//	}
//	author_id,err :=user.GetUserId(videoBody.Username)
//	if err!=nil{
//		common.JsonFail(w,500,err.Error())
//		return
//	}
//	video_info ,err:=AddVideo(author_id,videoBody.Video_name,)
//	if err!=nil{
//		common.JsonFail(w,500,err.Error())
//		return
//	}else{
//		common.JsonSucess(w,video_info,200,"ok")
//	}
//
//}

//func ListUserAllVideos(w http.ResponseWriter,r *http.Request,p httprouter.Params){
//	//lv := []map[string]string{}
//	ok,err:=user.VaildUser(w,r)
//	if !ok{
//		common.JsonSucess(w,"",200,err.Error())
//		return
//	}
//	username := p.ByName("usernmae")
//
//	s,err:=GetUserVideos(username)
//	if err!=nil{
//		log.Printf("huoqu user All videos error")
//		common.JsonFail(w,500,"huoqu shipin liebiao shibai ")
//		return
//	}
//	common.JsonSucess(w,s,200,"ok")
//
//}

func ListAllVideos(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//lv := []map[string]string{}
	ok,err:=user.VaildUser(w,r)
	if !ok{
		common.JsonSucess(w,"",200,err.Error())
		return
	}

	s,err:=GetAllVideos()
	if err!=nil{
		log.Printf("huoqu All videos error:%v",err)
		common.JsonFail(w,500,"huoqu shipin liebiao shibai ")
		return
	}
	common.JsonSucess(w,s,200,"ok")

}

type comment struct {
	Video_id string
	Author_id int
	Content string
}
func PostComment(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	ok,err:=user.VaildUser(w,r)
	if !ok{
		common.JsonSucess(w,"",200,err.Error())
		return
	}
	co := &comment{"",0,""}
	reqBody,_ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody,co);err!= nil{
		log.Println(err)
		common.JsonFail(w,500,"评论 失败")
		return
	}

	video_id := p.ByName("video_id")
	comment,err:=comments.AddComment(video_id,co.Author_id,co.Content)
	if err!=nil{
		log.Println(err)
		common.JsonFail(w,500,"评论 失败")
		return
	}else{
		common.JsonSucess(w,*comment,200,"ok")
	}





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

func GetUserVideos_sheqi(username string)([]map[string]string, error){
	var video_name  string
	var display_time string
	image_path := ""
	video_path := ""
	var m map[string]string
	var lv []map[string]string


	author_id ,err := user.GetUserId(username)
	if err!=nil{
		return nil,err
	}
	stmt,err := db.Db.Prepare("select video_name,dispaly_time,image_path,video_path from video_info where author_id=?")
	if  err!=nil{
		return nil,err
	}
	rows,err := stmt.Query(author_id)
	if  err!=nil{
		return nil,err
	}

	for rows.Next(){
		err:=rows.Scan(&video_name,&display_time,&image_path,&video_path)
		if err != nil{
			return nil,err
		}
		m["video_name"] = video_name
		m["display_time"] = display_time
		m["image_path"] = image_path
		m["video_path"] = video_path


		lv = append(lv,m)
		log.Println(lv)

	}
	return lv,nil

}

type video_info1 struct {
	Video_name string
	Author_id int
	Display_time string
	Image_path string
	Video_path string
}

type video_info2 struct {
	Author_id int
	Video_name string
	Username string
	Display_time string
	Image_path string
	Video_path string
}
func GetAllVideos()([]video_info2, error){
	//var video_name  string
	//var author_id int
	//var display_time string
	//var username string
	 //m :=map[string]string{}
	//image_path := ""
	//video_path := ""
	lv:= []video_info1{}
	lv2 := []video_info2{}
	video := video_info1{}
	video2 := video_info2{}


	stmt,err := db.Db.Prepare("select author_id,video_name,display_time,image_path,video_path from video_info")
	if  err!=nil{
		return nil,err
	}
	rows,err := stmt.Query()
	if  err!=nil{
		return nil,err
	}

	for rows.Next(){
		err:=rows.Scan(&video.Author_id,&video.Video_name,&video.Display_time,&video.Image_path,&video.Video_path)
		if err != nil{
			return nil,err
		}

		if err!=nil{
			return nil,err
		}
		//&video2.Username = &username
		//&video2.Video_path=&video.Video_path
		//&video2.Image_path=&video.Image_path
		//&video2.Display_time=&video.Display_time
		//&video2.Video_name=&video.Video_name
		//m["username"] = video.Author_id
		//m["video_name"] = video.Video_name
		//m["display_time"] = video.Display_time
		//m["image_path"] = video.Image_path
		//m["video_path"] = video.Video_path
		lv = append(lv,video)

	}
	//var s int
	for i := range lv{
		func(s int){
			username,_ := user.GetUsername(lv[s].Author_id)
			video2.Username = username
			video2.Author_id=lv[s].Author_id
			video2.Video_path=lv[s].Video_path
			video2.Image_path=lv[s].Image_path
			video2.Display_time=lv[s].Display_time
			video2.Video_name=lv[s].Video_name
			lv2 = append(lv2,video2)
		}(i)

	}
	//log.Println(lv)
	return lv2,nil
}
func GetUserVideos(aid int)([]video_info2, error){
	//var video_name  string
	//var author_id int
	//var display_time string
	//var username string
	//m :=map[string]string{}
	//image_path := ""
	//video_path := ""
	lv:= []video_info1{}
	lv2 := []video_info2{}
	video := video_info1{}
	video2 := video_info2{}


	stmt,err := db.Db.Prepare("select author_id,video_name,display_time,image_path,video_path from video_info where author_id = ?")
	if  err!=nil{
		return nil,err
	}
	rows,err := stmt.Query(aid)
	if  err!=nil{
		return nil,err
	}

	for rows.Next(){
		err:=rows.Scan(&video.Author_id,&video.Video_name,&video.Display_time,&video.Image_path,&video.Video_path)
		if err != nil{
			return nil,err
		}

		if err!=nil{
			return nil,err
		}
		//&video2.Username = &username
		//&video2.Video_path=&video.Video_path
		//&video2.Image_path=&video.Image_path
		//&video2.Display_time=&video.Display_time
		//&video2.Video_name=&video.Video_name
		//m["username"] = video.Author_id
		//m["video_name"] = video.Video_name
		//m["display_time"] = video.Display_time
		//m["image_path"] = video.Image_path
		//m["video_path"] = video.Video_path
		lv = append(lv,video)

	}
	//var s int
	for i := range lv{
		func(s int){
			username,_ := user.GetUsername(lv[s].Author_id)
			video2.Username = username
			video2.Author_id=lv[s].Author_id
			video2.Video_path=lv[s].Video_path
			video2.Image_path=lv[s].Image_path
			video2.Display_time=lv[s].Display_time
			video2.Video_name=lv[s].Video_name
			lv2 = append(lv2,video2)
		}(i)

	}
	//log.Println(lv)
	return lv2,nil
}