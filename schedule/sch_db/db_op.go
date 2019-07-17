package sch_db

import (
	"log"
	"sis_video_go/db"
)

func ReadDelVideoRecord(count int)([]string,error){
	stmt ,err :=db.Db.Prepare("select video_id from del_Video_Record limit ?")
	var video_ids []string
	if err != nil{
		log.Println(err)
		return video_ids,err
	}
	rows,err := stmt.Query(count)
	if err != nil{
		log.Println(err)
		return video_ids,err
	}
	var id string
	for rows.Next(){
		if err := rows.Scan(&id);err != nil{
			log.Println(err)
			return video_ids,err
		}
		video_ids = append(video_ids,id)


	}
	defer stmt.Close()
	return video_ids,nil
}


func DelVideoDelRecord(vid string)error{
	stmt ,err:= db.Db.Prepare("delete from del_Video_Record where video_id=?")
	if err !=nil{
		log.Println(err)
		return err
	}
	_,err = stmt.Exec(vid)
	if err != nil{
		log.Println(err)
		return err
	}

	stmt.Close()
	return nil
}