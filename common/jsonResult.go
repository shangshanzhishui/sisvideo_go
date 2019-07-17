package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type JsonRes struct {
	Msg   string `json:"msg"`
	Response string `json:"response"`
	Data     interface{} `json:"data"`
	Error    int `json:"error"`
	Nums     int `json:"nums"`
}

var jsonSucess = JsonRes{
	Msg: "",
	Response:"ok",
	Data:0,
	Error:0,
	Nums:0,

}

var jsonFail = JsonRes{
	Msg: "",
	Response:"fail",
	Data:0,
	Error:0,
	Nums:0,

}
func JsonSucess(w http.ResponseWriter,v interface{},code int,msg string) error{
	jsonSucess.Data = v
	jsonSucess.Msg = msg
	w.WriteHeader(code)
	data ,err := json.Marshal(jsonSucess)
	if err != nil{
		return err
	}
	//w.WriteHeader(code)
	w.Write(data)
	return nil
}

func JsonFail(w http.ResponseWriter,code int,msg string) error{

	jsonFail.Msg = msg
	w.WriteHeader(code)
	data ,err := json.Marshal(jsonFail)
	if err != nil{
		return err
	}
	//w.WriteHeader(code)
	w.Write(data)
	return nil
}

func UnMarshalHson(r *http.Request,v interface{}) error{
	result ,err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}
	json.Unmarshal([]byte(bytes.NewBuffer(result).String()),v)
	return nil
}
