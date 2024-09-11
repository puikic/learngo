package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func get() {
	resp, err := http.Get("http://localhost:5656/boy")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //一定用Close，否则会协程泄露
	io.Copy(os.Stdout, resp.Body)
	for k, v := range resp.Header { //打印response header
		fmt.Printf("%s = %v\n", k, v)
	}
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
}

func post() {
	reader := strings.NewReader("hello server")
	resp, err := http.Post("http://localhost:5656/boy", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //一定用Close，否则会协程泄露
	io.Copy(os.Stdout, resp.Body)
	for k, v := range resp.Header { //打印response header
		fmt.Printf("%s = %v\n", k, v)
	}
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
}

func complexHttpRequest() {
	reader := strings.NewReader("hello server")
	// http.NewRequest("POST", "http://localhost:5656/boy", reader)
	if req, err := http.NewRequest("POST", "http://localhost:5656/boy", reader); err != nil {
		panic(err)
	} else {
		//自定义请求头
		req.Header.Add("User-Agent", "中国")
		req.Header.Add("MyHeaderKey", "MyHeaderValue")
		//自定义cookie
		req.AddCookie(&http.Cookie{
			Name:    "auth",
			Value:   "passwd",
			Path:    "/",
			Domain:  "localhost",
			Expires: time.Now().Add(time.Duration(time.Hour)),
		})
		client := &http.Client{
			Timeout: 100 * time.Millisecond,
		}
		if resp, err := client.Do(req); err != nil { //提交http请求
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			io.Copy(os.Stdout, resp.Body)
			for k, v := range resp.Header { //打印response header
				fmt.Printf("%s = %v\n", k, v)
			}
			fmt.Println(resp.Proto)
			fmt.Println(resp.Status)
		}
	}

}
func main() {
	// get()
	// post()
	complexHttpRequest()
}
