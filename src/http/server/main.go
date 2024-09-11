package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func BoyHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Body)
	io.Copy(os.Stdout, r.Body)
	for k, v := range r.Header { //打印request header
		fmt.Printf("%s = %v\n", k, v)
	}
	for _, cookie := range r.Cookies() {
		fmt.Printf("%s = %s\n", cookie.Name, cookie.Value)
	}
	fmt.Fprint(w, "Hello Boy!")
}
func GirlHandler(w http.ResponseWriter, r *http.Request) {
	io.Copy(os.Stdout, r.Body)
	for k, v := range r.Header { //打印request header
		fmt.Printf("%s = %v\n", k, v)
	}
	fmt.Fprint(w, "Hello Girl!")
}

func main() {
	//定义路由
	http.HandleFunc("/boy", BoyHandler)
	http.HandleFunc("/girl", GirlHandler)
	//把服务启起来
	http.ListenAndServe(":5656", nil) //如果不发生error,会一直阻塞
}
