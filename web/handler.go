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
	"sis_video_go/api/comments"
	"sis_video_go/api/user"
	"sis_video_go/api/video"
	"sis_video_go/common"
)

type Home struct {
	Name string
}

type User struct {
	Name string
	C string
}

func MyHomeHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//video.ListAllVideos(w,r,p)
	username := p.ByName("username")
	id,err := user.GetUserId(username)
	s,err:=video.GetUserVideos(id)
	if err != nil{
		log.Printf("parsing template home.html error:%s:",err)
		common.JsonFail(w,200,"用户名错误")
		return
	}
	//sid ,err := r.Cookie("session")
	log.Println(s)
	if err != nil{
		log.Println(err)
	}
	t,err := template.ParseFiles("./templates/home.html")
	if err != nil{
		log.Printf("parsing template home.html error:%s:",err)
		common.JsonFail(w,500,"chucuo")
		return
	}
	t.Execute(w,s)
	return
	//if sid.Value =="ds"{
	//	http.Redirect(w,r,"/user",http.StatusFound)
	//}

}


func RegiserHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){

		p1 := &Home{Name:"SISIPHUS"}
		t,err := template.ParseFiles("./templates/register.html")
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
	username := p.ByName("username")
	url := "/video/"+username
	t,_:=template.ParseFiles("./templates/upload.html")
	t.Execute(w,url)



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
type Res struct {
	Aid int
	Videos interface{}
}
func indexHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//video.ListAllVideos(w,r,p)
	s,err:=video.GetAllVideos()
	//username ,err := r.Cookie("username")
	//id,err:=user.GetUserId(username.Value)
	log.Println(s)
	if err != nil{
		log.Println(err)
		return
	}
	//user.GetUserId(username.Value)
	log.Println(s)
	t,err := template.ParseFiles("./templates/index.html")
	if err != nil{
		log.Printf("parsing template home.html error:%s:",err)
		return
	}
	res := Res{}
	res.Aid = 1
	res.Videos = s
	t.Execute(w,res)
	return
	//if sid.Value =="ds"{
	//	http.Redirect(w,r,"/user",http.StatusFound)
	//}

}
type res struct {
	Comment_url string
	Video_url string
	Comments interface{}
}
func videoHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	video_id := p.ByName("video_id")
	re := &res{}
	comments,err := comments.GetComments(video_id)
	re.Comments = comments
		log.Printf("pinglun:%v",comments)
	if comments==nil{
		re.Comments = nil
	}
		t,err := template.ParseFiles("./templates/comment.html")
		if err != nil{
			log.Printf("parsing template comment.html error:%s:",err)
			return
		}
	url := "/videos/"+video_id+"/comment"

	re.Comment_url= url
	re.Video_url = "/video/"+video_id
	log.Println(re.Video_url)
	t.Execute(w,re)
}



//func videoCommentHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
//
//}

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

	proxy.ServeHTTP(w,r)

}