package main

import (
	"net/http"
	"myGo/src/database"
	"log"
)

func index(resp http.ResponseWriter, req *http.Request) {

}
func login(resp http.ResponseWriter, req *http.Request) {

}
func err(resp http.ResponseWriter, req *http.Request) {

}

func authenticate(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	user, err := database.UserByEmail(req.PostFormValue("email"))
	if err != nil {
		log.Fatal(err)
		return
	}
	if user.Password != database.Encrypt(req.PostFormValue("password")) {
		// 密码不正确，返回302
		http.Redirect(resp, req, "/login", 302)
	}
	if user.Password != database.Encrypt(req.PostFormValue("password")) {
		// 密码正确，存储session至cookie
		sesson:=database.Session{}//todo 这里应该有个根据user创建session的方法
		cookie:=http.Cookie{
			Name: "_cookie",
			Value:sesson.Uuid,
			HttpOnly:true,
		}
		// Q:这里为什么要传入resp参数
		http.SetCookie(resp,&cookie)
		http.Redirect(resp,req,"/",302)
	}
}
