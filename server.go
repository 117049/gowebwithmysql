package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method=", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("D:\\GoCode\\src\\gowebwithmysql\\gowebwithmysql\\login.tpl")
		log.Println(t.Execute(w, nil))
	}else {
		_ = r.ParseForm()
		fmt.Println("username=", r.Form["username"])
		fmt.Println("password=", r.Form["password"])
		if pwd := r.Form.Get("password"); pwd == "12345" {
			fmt.Fprintf(w, "欢迎登录, Hello 帅气的%s!", r.Form.Get("username"))
		}else {
			fmt.Fprintf(w, "密码错误，请重新输入")
		}
	}
}

func SayHelloName (w http.ResponseWriter, r *http.Request)  {
	_ = r.ParseForm() // 解析参数
	fmt.Println(r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Host: ", r.Host)
	for k, v := range r.Form{
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello acho!")
}

func main() {
	http.HandleFunc("/", SayHelloName) // 设置访问的路由
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("ListenAndServer: ", err)
	}
}






