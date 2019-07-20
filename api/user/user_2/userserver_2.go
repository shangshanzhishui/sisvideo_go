package user_2

import (
	"log"
	"sis_video_go/db"
)

func GetUserId(username string) (int,error){
	id :=0
	stmt ,err := db.Db.Prepare("select id from user where username = ?")
	if err != nil{
		log.Println(err)
		return 0,err
	}
	err = stmt.QueryRow(username).Scan(&id)
	if err != nil{
		return 0,err
	}
	defer stmt.Close()
	return id,nil
}
