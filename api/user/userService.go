package user
import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
  	"log"
	"net/http"
	"sis_video_go/db"
	session2 "sis_video_go/session"
)

type User struct {
	Username   string `json:"username"`
	Pwd        string `json:"pwd"`
	Login_time string `json:"login_time"`
}

func AddUser(usernam string,pw string) error {
	fmt.Print(usernam,pw)
	stmt,err := db.Db.Prepare("insert into user (username,pwd) values (?,?)")
	if err != nil{
		fmt.Println(err)
		return err
	}
	_,err = stmt.Exec("nihao","123")
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
	session,err := r.Cookie("session")
	if err != nil{
		log.Println(err)
		return false,err
	}

	session_id := session.Value
	_,ok := session2.SessionMap.Load(session_id)
	if ok{
		return true,nil
	}else{
		err := errors.New("wei denglu")
		return false,err
	}
}