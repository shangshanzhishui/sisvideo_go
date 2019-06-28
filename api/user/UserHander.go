package user
import (
	_ "github.com/go-sql-driver/mysql"
  	"log"
	"sis_video_go/db"
)

func Login(username string,pwd string) error {
	stmt,err := db.Db.Prepare("insert into user (username,pwd) values (?,?)")
	if err != nil{

		return err
	}
	stmt.Exec(username,pwd)
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
	stmt.QueryRow(username).Scan(&pwd)
	defer stmt.Close()
	return pwd ,nil
}

func DeleteUser(username string,pwd string ) error{
	stmt , err := db.Db.Prepare("delete from users where username=? and pwd =?")
	if err != nil{
		log.Println(err)
		return nil
	}
	stmt.Exec(username,pwd)
	defer stmt.Close()
	return nil

}