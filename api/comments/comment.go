package comments

import (
	"database/sql"
	"log"
	"sis_video_go/db"
	"sis_video_go/utils"
	"time"
)

type Comment struct {
	Id           string `json:"id"`
	Video_id     string `json:"video_id"`
	Author_id    int    `json:"author_id"`
	Content      string  `json:"content"`
	Display_time string  `json:"display_time"`
}

func AddComment(vid string,aid int,content string) (*Comment,error){
	id := utils.NewUUID()
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


type Comment2 struct {
	Id           string `json:"id"`
	Author_name  string `json:"author_name"`
	Content      string `json:"content"`
	Display_time string `json:"display_time"`
}
func GetComments(vid string)([]Comment2,error){

	comments := Comment2{}
	var res []Comment2
	stmt,err := db.Db.Prepare("select comments.id ,user.username,comments.content,display_time from comments inner join user on comments.author_id=user.id where comments.video_id = ?")
	//"select comments.id ,user.username,comments.content,display_time from comments inner join user on comments.author_id=user.id where comments.video_id = ?"
	//log.Printf("shujulianjie:%v",stmt)
	if stmt ==nil{
		return nil,nil
	}
	rows,err := stmt.Query(vid)
	if err != nil{
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&comments.Id,&comments.Author_name,&comments.Content,&comments.Display_time); err!= nil{
			return res,err
		}
		res  = append(res,comments)
	}
	defer stmt.Close()
	if len(res) ==0{
		return nil,nil
	}
	log.Println("123")
	log.Println(res)
	return res,nil
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