package main

import (
	"fmt"
	"net/http"
)

//注释
//建立一个http服务, 增加一个http登录
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "login")
	})
	http.ListenAndServe(":8080", nil)
}
