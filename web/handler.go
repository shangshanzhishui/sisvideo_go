package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sis_video_go/api/user"
	"sis_video_go/common"
)

type Home struct {
	Name string
}

type User struct {
	Name string
}

func homeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	sid ,err := r.Cookie("session")
	if err != nil{
		p1 := &Home{Name:"SISIPHUS"}
		t,err := template.ParseFiles("./templates/home.html")
		if err != nil{
			log.Printf("parsing template home.html error:%s:",err)
			return
		}
		t.Execute(w,p1)
		return
	}
	if sid.Value =="ds"{
		http.Redirect(w,r,"/user",http.StatusFound)
	}

}


func userHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	sid ,err := r.Cookie("session")
	if err != nil{
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}

	var u *User
	username := r.FormValue("username")
	username ,err = user.GetUserBySessionId(sid.Value)
	if err != nil{
		log.Printf("parsing user.html error:%v",err)
		return
	}
	u= &User{Name:username}
	t,err := template.ParseFiles("./templates/user.html")
	t.Execute(w,u)

}


func apiHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	if r.Method != http.MethodPost{
		common.JsonFail(w,400,"qingqiufangshi must is POST")
	}
	res,_ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res,apibody);err !=nil{
		common.JsonFail(w,500,"error")
		return
	}
	request(apibody,w,r)

}