package main

import (
	"FileShare/src/AppSet"
	"FileShare/src/Delete"
	"FileShare/src/Download"
	"FileShare/src/Index"
	"FileShare/src/Upload"
	"embed"
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

//go:embed Static
var static embed.FS

//命令行参数
var n = flag.String("n", "FileShare", "分享站名字")
var d = flag.String("d", "./Data", "数据存放地址")
var p = flag.Int("p", 8080, "监听端口号")

//启动http服务器
func start(p *int) {
	//启动服务
	err := http.ListenAndServe(":"+strconv.Itoa(*p), nil)
	if err != nil {
		fmt.Println("服务启动失败：" + err.Error())
	}
}

//主函数
func main() {
	//解析命令行参数
	flag.Parse()
	//个人设置
	AppSet.SetName(*n)
	AppSet.SetData(*d)
	//打印信息
	fmt.Println("文件分享站：" + *n)
	fmt.Println("数据地址：" + *d)
	fmt.Println("监听端口：" + strconv.Itoa(*p))
	fmt.Println("启动服务...")
	//设置监听
	staticHandle := http.FileServer(http.FS(static))
	http.Handle("/Static/js/", staticHandle)
	http.Handle("/Static/css/", staticHandle)
	http.Handle("/Static/img/", staticHandle)
	http.HandleFunc("/", Index.Index)
	http.HandleFunc("/Upload", Upload.Upload)
	http.HandleFunc("/Download", Download.Download)
	http.HandleFunc("/Delete", Delete.Delete)
	//启动http服务器，开始监听
	start(p)
}
