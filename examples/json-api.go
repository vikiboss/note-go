package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	IP   = "127.0.0.1"
	Port = "8080"
)

// 定义接口返回的 JSON 数据结构
type message struct {
	Message string `json:"message"`
}

func main() {
	// 注册处理请求的函数
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/hello", helloHandler)

	// 启动服务器并监听指定地址
	addr := fmt.Sprintf("%s:%s", IP, Port)
	log.Printf("Server listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// 处理 /hello 请求的函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 从 URL 参数中获取 name
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	// 构造返回的 JSON 结果
	m := message{Message: fmt.Sprintf("Hello, %s!", name)}
	result, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置 HTTP 响应头
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// 返回 JSON 结果
	w.Write(result)
}
