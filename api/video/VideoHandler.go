package video

import (
	"database/sql"
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
	id := utils.CreateUUID()
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