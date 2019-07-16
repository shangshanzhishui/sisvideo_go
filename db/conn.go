package db

import (
	"database/sql"
	"fmt"
	"log"
	"sis_video_go/setting"
	_"github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init() {

	dataSourceName := setting.Mysql_db_name +":"+ setting.Mysql_Password +"@(localhost:3306)/sis_video?charset=utf8"
	fmt.Println(dataSourceName)
	Db,err = sql.Open("mysql",dataSourceName)
	if err != nil{
		panic(err)
	}
	if err := Db.Ping(); err != nil{
		panic(err)
	}
	log.Printf("db is init %s",Db)

}