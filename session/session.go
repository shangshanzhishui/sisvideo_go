package session

import (
	"database/sql"
	"log"
	"sis_video_go/db"
	"sis_video_go/session/sessionServer"
	"sis_video_go/setting"
	"sis_video_go/utils"
	"strconv"
	"time"
)

type Session struct {
	Sesion_id string
	Author_id int
	TTL       string
}

type Session2 struct {
	Sesion_id string
	Author_id int
	TTL       int64
}



func CreateSessionId(author_id int) string{
	id := utils.NewUUID()
	ctime := time.Now().UnixNano() / 1000000
	ttl := strconv.FormatInt(ctime + setting.SessionExpired,10)
	session := &Session{Sesion_id:id,Author_id:author_id,TTL:ttl}
	sessionServer.SessionMap.Store(id,session)

	AddSession(id,ttl,author_id)
	return id
}





func SessionIsExpired(id string)bool{
	session,ok := sessionServer.SessionMap.Load(id)
	if ok{
		ctime := time.Now().UnixNano() / 1000000
		ttl,_ := strconv.ParseInt(session.(*Session).TTL,10,64)
		if ttl< ctime{
			return true
		}
	}
	DeleteSession(id)
	return false
}

func DeleteSession(sid string){
	sessionServer.SessionMap.Delete(sid)
	DeleteSessionDB(sid)
}



func AddSession(sid string,ttl string,author_id int) error{

	stmt,err := db.Db.Prepare("insert into session (session_id,author_id,TTL) values (?,?,?)")
	if err != nil{
		return err
	}
	_,err = stmt.Exec(sid,ttl,author_id)
	if err != nil{
		return err
	}
	defer stmt.Close()
	return nil
}


func FixSession(sid string)(*Session,error){
	stmt,err := db.Db.Prepare("select TTL ,author_id from sessions where session_id =?")
	session := &Session{}
	if err != nil{
		return nil,err
	}
	err = stmt.QueryRow(sid).Scan(session.TTL,session.Author_id)
	if err != nil && err != sql.ErrNoRows{
		return nil,err
	}
	defer stmt.Close()

	return session,nil
}

func DeleteSessionDB(sid string)error{
	stmt ,err := db.Db.Prepare("delete from session where session_id = ?")
	if err != nil{
		log.Printf("%s",err)
		return nil

	}

	if _,err := stmt.Query(sid);err != nil{
		return err
	}
	return nil
}


func LoadAllSession(){
	//m := &sync.Map{}
	session2 := &Session2{}
	session := &Session{}
	stmt,err := db.Db.Prepare("select * from sessions")
	if err != nil{

		log.Println(err)
	}

	rows,err := stmt.Query()
	if err != nil{

		log.Println(err)
	}

	for rows.Next(){
		err := rows.Scan(session.Sesion_id,session.Author_id,session.TTL)
		if err != nil{

			log.Println(err)
			break
		}
		ttl,err := strconv.ParseInt(session.TTL,10,64)
		if err != nil{

			log.Println(err)
			break
		}
		session2.Author_id = session.Author_id
		session2.Sesion_id = session.Sesion_id
		session2.TTL = ttl
		sessionServer.SessionMap.Store(session2.Sesion_id,session2)
		log.Printf("session_id:%v,ttl:%v",session2.Sesion_id,ttl)
	}
	return


}
