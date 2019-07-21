package user
import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
  	"log"
	"net/http"
	"sis_video_go/db"
	"sis_video_go/session/sessionServer"
)

type User struct {
	Username   string `json:"username"`
	Pwd        string `json:"pwd"`
	Login_time string `json:"login_time"`
}

func AddUser(username string,pwd string) error {
	fmt.Print(username,pwd)
	stmt,err := db.Db.Prepare("insert into user (username,pwd) values (?,?)")
	if err != nil{
		fmt.Println(err)
		return err
	}
	_,err = stmt.Exec(username,pwd)
	if err != nil{
		return nil
	}
	defer stmt.Close()
	return nil
}

func GetUser(username string) (string,error){
	stmt,err := db.Db.Prepare("select pwd from user where username = ?")
	if err != nil{
		log.Printf("s%",err)
		return "",err
	}
	var pwd string
	err = stmt.QueryRow(username).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{
		return "",err

	}
	defer stmt.Close()
	return pwd ,nil
}

func DeleteUser(username string,pwd string ) error{
	stmt , err := db.Db.Prepare("delete from user where username=? and pwd =?")
	if err != nil{
		log.Println(err)
		return err
	}
	_,err = stmt.Exec(username,pwd)
	if err != nil{
		return err
	}
	defer stmt.Close()
	return nil

}

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

func GetUsername(id int) (string,error){
	var username string
	stmt ,err := db.Db.Prepare("select username from user where id = ?")
	if err != nil{
		log.Println(err)
		return "",err
	}
	err = stmt.QueryRow(id).Scan(&username)
	if err != nil{
		return "",err
	}
	defer stmt.Close()
	return username,nil
}

func GetUserBySessionId(sid string) (string,error ){
	stmt,err := db.Db.Prepare("select author_id from sessions where session_id is ?")
	if err != nil{
		log.Println(err)
		return "",err
	}
	author_id := 0
	err = stmt.QueryRow(sid).Scan(&author_id)
	if err != nil{
		return "",err
	}

	//查询user表
	stmt1,err := db.Db.Prepare("select username from user where id is ?")
	if err != nil{
		log.Println(err)
		return "",err
	}
	username := ""
	err = stmt1.QueryRow(author_id).Scan(&username)
	if err != nil{
		return "",err
	}
	defer stmt1.Close()
	return username,nil
}

func VaildUser(w http.ResponseWriter,r *http.Request) (bool,error){
	log.Println("jinru")
	session,err := r.Cookie("session")
	if err != nil{
		log.Println(err)
		return false,err
	}

	session_id := session.Value
	log.Println(session_id)
	log.Println(sessionServer.SessionMap)
	_,ok := sessionServer.SessionMap.Load(session_id)
	log.Println(ok)
	if ok{
		return true,nil
	}else{
		err := errors.New("wei denglu")
		return true,err
	}
}