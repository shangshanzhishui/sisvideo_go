package user
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
  	"log"
	"sis_video_go/db"
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