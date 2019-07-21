package sessionServer

import (
	"log"
	"sis_video_go/api/user/user_2"
	"sis_video_go/db"
	"sis_video_go/setting"
	"sis_video_go/utils"
	"strconv"
	"sync"
	"time"
)

var SessionMap *sync.Map
func init(){
	SessionMap = &sync.Map{}

}

type Session struct {
	Sesion_id string
	Author_id int
	TTL       string
}

func CreateSessionId(author_id int) string{
	id := utils.NewUUID()
	ctime := time.Now().UnixNano() / 1000000
	ttl := strconv.FormatInt(ctime + setting.SessionExpired,10)
	session := &Session{Sesion_id:id,Author_id:author_id,TTL:ttl}
	SessionMap.Store(id,session)

	err := AddSession(id,ttl,author_id)
	if err!= nil{
		log.Println(err)
	}
	return id
}

func AddSession(sid string,ttl string,author_id int) error{

	stmt,err := db.Db.Prepare("insert into sessions (session_id,author_id,TTL) values (?,?,?)")
	if err != nil{
		return err
	}
	_,err = stmt.Exec(sid,author_id,ttl)
	if err != nil{
		return err
	}
	defer stmt.Close()
	return nil
}

func CreateSessionIdByUname(Username string) (string,error){
	id := utils.NewUUID()
	ctime := time.Now().UnixNano() / 1000000
	ttl := strconv.FormatInt(ctime + setting.SessionExpired,10)
	author_id,err := user_2.GetUserId(Username)
	if err != nil{
		return "",err
	}
	session := &Session{Sesion_id:id,Author_id:author_id,TTL:ttl}
	SessionMap.Store(id,session)
	AddSession(id,ttl,author_id)
	return id,nil
}