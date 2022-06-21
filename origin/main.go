package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func SayHello (w http.ResponseWriter, r *http.Request)  {
	_ = r.ParseForm() // 解析参数
	fmt.Println(r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Host: ", r.Host)
	for k, v := range r.Form{
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello web, %s!", r.Form.Get("name"))
}

func main() {
	http.HandleFunc("/", SayHello) // 设置访问的路由
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("ListenAndServer: ", err)
	}
}
