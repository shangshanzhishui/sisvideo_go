package user
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
  	"log"
	"sis_video_go/db"
)

type User struct {
	Username   string `json:"username"`
	Pwd        string `json:"pwd"`
	Login_time string `json:"login_time"`
}

func AddUser(username string,pwd string) error {
	stmt,err := db.Db.Prepare("insert into user (username,pwd) values (?,?)")
	if err != nil{

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
	stmt , err := db.Db.Prepare("delete from users where username=? and pwd =?")
	if err != nil{
		log.Println(err)
		return nil
	}
	_,err = stmt.Exec(username,pwd)
	if err != nil{
		return err
	}
	defer stmt.Close()
	return nil

}