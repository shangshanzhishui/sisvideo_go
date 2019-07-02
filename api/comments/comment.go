package comments

import (
	"database/sql"
	"sis_video_go/db"
	"sis_video_go/utils"
	"time"
)

type Comment struct {
	Id           string
	Video_id     string
	Author_id    int
	Content      string
	Display_time string
}

func AddComment(vid string,aid int,content string) (*Comment,error){
	id := utils.CreateUUID()
	t := time.Now()
	ctime := t.Format("jan 02 2006,15:04:05")
	stmt,err := db.Db.Prepare("insert into comments (id ,video_id,author_id,content,display_time) values (?,?,?,?,?)")
	if err != nil{
		return nil,err
	}
	_,err = stmt.Exec(id,vid,aid,content,ctime)
	if err != nil{
		return nil,err
	}
	defer stmt.Close()
	res := &Comment{id,vid,aid,content,ctime}
	return res,nil
}


func GetComment(id string) (*Comment,error){
	comment := &Comment{}
	stmt,err := db.Db.Prepare("select id,cideo_id,author_id,content,dispaly_time from comments where id=?")
	if err != nil{
		return nil,err
	}

	err = stmt.QueryRow(id).Scan(comment)
	if err!= nil && err != sql.ErrNoRows{
		return nil,err
	}
	if err == sql.ErrNoRows{
		return nil,nil
	}
	defer stmt.Close()
	return comment,nil

}

func DeleteComment(id string) error{
	stmt,err := db.Db.Prepare("delete from comments where id =?")
	if err!=nil{
		return err
	}

	_,err = stmt.Exec(id)
	if err != nil{
		return err
	}
	defer stmt.Close()
	return nil
}