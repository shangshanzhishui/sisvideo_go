package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"sis_video_go/common"
)

type ApiBody struct {
	Url 	string  `json:"url"`
	Method  string  `json:"method"`
	body    string  `json:"body"`

}

var client *http.Client

func init(){
	client = &http.Client{}
}

func  request( b *ApiBody,w http.ResponseWriter,r *http.Request){
	var resp *http.Response
	var err error
	switch b.Method {
	case http.MethodGet:
		req,_ := http.NewRequest("GET",b.Url,nil)
		req.Header = r.Header
		resp,err = client.Do(req)
		if err != nil{
			log.Println(err)
			return
		}
		Myresponse(w,resp)
	case http.MethodPost:

		req,_ := http.NewRequest("POST",b.Url,nil)
		req.Header = r.Header
		resp,err = client.Do(req)
		if err != nil{
			log.Println(err)
			return
		}
		Myresponse(w,resp)
	case http.MethodDelete:
		req,_ := http.NewRequest("DELETE",b.Url,bytes.NewBuffer([]byte(b.body)))
		req.Header = r.Header
		resp,err = client.Do(req)
		if err != nil{
			log.Println(err)
			return
		}
		Myresponse(w,resp)
	default:


	}
}

func Myresponse(w http.ResponseWriter,r *http.Response){
	res,err := ioutil.ReadAll(r.Body)
	if err != nil{
		common.JsonFail(w,500,"error")
		return
	}
	common.JsonSucess(w,string(res),r.StatusCode,"")
}