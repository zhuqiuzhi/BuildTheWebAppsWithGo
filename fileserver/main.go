package main

import (
	"net/http"
)

func main() {
	// http.FileServer的输入参数为接口 http.FileSystem, 返回值为满足接口 http.Handler 的隐藏类型fileHandler
	// 将当前目录强制转换为 http.Dir, http.Dir 接收文件目录路径(如当前目录 "." , 上一级目录 "..")
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}