package db

import (
	"database/sql"
	"sis_video_go/setting"
)

var (
	Db *sql.DB
	err error
)

func init() {

	dataSourceName := setting.Mysql_db_name + setting.Mysql_Password +"@#@(127.0.0.1:3306)/sis_video?charset=utf8"
	Db,err := sql.Open("mysql",dataSourceName)
	if err != nil{
		panic(err.Error())
	}

}