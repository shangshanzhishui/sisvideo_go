package setting

type Error struct {
	ErrorMsg string `json:"error_msg"`
	ErrCode string `json:"err_code"`
}



type ErrHttpResponse struct {
	HttpResCode int
	Error Error
}

var Mysql_db_name ="root"
var Mysql_Password = "12345678"
var SessionExpired = int64(30 * 60 *1000)

const (
	VIDEO_DICTORY = "./../videos" //shipinlujing
)