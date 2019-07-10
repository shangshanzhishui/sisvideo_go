package setting

type Error struct {
	Error string `json:"error"`
	ErrCode string `json:"err_code"`
}

type ErrHttpResponse struct {
	HttpResCode int
	Error Error
}

var Mysql_db_name ="root"
var Mysql_Password = "12345678"
var SessionExpired = int64(30 * 60 *1000)