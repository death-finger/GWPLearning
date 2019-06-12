package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 打印启动信息
	fmt.Println("ChitChat", version(), "started at", config.Address)

	// 创建一个多路复用器
	mux := http.NewServeMux()
	// 处理静态文件
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// 路由处理
	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

}
