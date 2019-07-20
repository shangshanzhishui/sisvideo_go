package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sis_video_go/api/user"
	"sis_video_go/common"
)

type Home struct {
	Name string
}

type User struct {
	Name string
	C string
}

func indexHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	sid ,err := r.Cookie("session")
	if err != nil{
		p1 := &Home{Name:"SISIPHUS"}
		t,err := template.ParseFiles("./templates/index.html")
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
func RegiserHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){

		p1 := &Home{Name:"SISIPHUS"}
		t,err := template.ParseFiles("./templates/re.html")
		if err != nil{
			log.Printf("parsing template home.html error:%s:",err)
			return
		}
		t.Execute(w,p1)
		return
	}
func loginHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){

		p1 := &Home{Name:"SISIPHUS"}
		t,err := template.ParseFiles("./templates/login.html")
		if err != nil{
			log.Printf("parsing template home.html error:%s:",err)
			return
		}
		t.Execute(w,p1)
		return
	}



func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	username,err:=r.Cookie("username")
	url := "/video/"+username.Value

	t,err := template.ParseFiles("./templates/upload.html")
	if err != nil{
		log.Printf("parsing template home.html error:%s:",err)
		return
	}
	t.Execute(w,url)
	//common.JsonSucess(w,`{"username1":"123"}`,200,"ok")

	return
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
	t,err := template.ParseFiles("./templates/index.html")
	if err!=nil{
		log.Println(err)
	}

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

func proxyScheHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	u,_ := url.Parse("http://127.0.0.1:9003/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w,r)
}


func proxyServeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	u,_ := url.Parse("http://127.0.0.1:9001/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w,r)
}

func proxyStreamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	u,err := url.Parse("http://127.0.0.1:9002/")
	log.Println(u)
	if err !=nil{
		log.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	log.Println(*proxy)
	proxy.ServeHTTP(w,r)

}