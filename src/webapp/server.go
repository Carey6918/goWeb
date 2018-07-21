package main

import (
	"net/http"
	"fmt"
)

/*
 * 处理器函数：从request提取相关信息，创建一个HTTP响应，再从ResponseWriter接口将响应返回给客户端
 * writer: ResponseWriter接口	request：指向request结构的指针
 */
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world,%s!", request.URL.Path[1:])
}

func main() {
	// 创建一个默认的多路复用器
	mux := http.NewServeMux()
	// 将发送至根URL的请求重定向到服务器--->重定向到名为index的处理器函数
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/err", err)

	// 自定义使用mux作为多路复用器的服务器
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	// 启动并监听服务器
	server.ListenAndServe()
}
