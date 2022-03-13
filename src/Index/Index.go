package Index

import (
	"FileShare/src/AppSet"
	"FileShare/src/SpaceDate"
	_ "embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"text/template"
)

//go:embed index.html
var index string

//go:embed welcome.html
var welcome string

var Iserr = 0

//获取文章
func Index(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	//获取数据文件夹目录和分享站名字
	FilePath := ""
	if len(request.Form["dir"]) == 0 {
		Name := AppSet.GetName()
		//返回数据
		templateHtml2(Name, response)
	} else {
		FilePath = request.Form["dir"][0]
		//创建数据文件夹
		err := os.MkdirAll(AppSet.GetData()+"/"+FilePath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		//更新时间
		SpaceDate.UpDate(FilePath)
		//扫描文件夹
		Files := scanFiles(FilePath)
		Name := AppSet.GetName()
		if Iserr == 1 {
			Iserr = 0
			http.Redirect(response, request, "./", http.StatusFound)
		}
		//返回数据
		templateHtml1(Files, Name, response)
	}
}

//扫描文件夹下的所有文件，返回文件名的集合
func scanFiles(FilePath string) string {
	var names string
	files, err := ioutil.ReadDir(AppSet.GetData() + "/" + FilePath)
	if err != nil {
		fmt.Println(err)
		Iserr = 1
		return ""
	}
	for _, f := range files {
		prefix := path.Ext(f.Name())
		if prefix == "" {
			prefix = "file"
		} else {
			prefix = prefix[1:]
		}
		if f.IsDir() {
			names += "<div onclick='ToDir(\"" + f.Name() + "\")' class='item folder'>" +
				"			<img src='./Static/img/icons/files.svg'>" +
				"			<span>" + f.Name() + "</span>" +
				"			<div class='delete' onclick='Delete(\"" + FilePath + "/" + f.Name() + "\")'>删除</div>" +
				"		</div>"
		} else {
			names += "<div class='item'>" +
				"		<img src='./Static/img/icons/" + prefix + ".svg'>" +
				"		<span>" + f.Name() + "</span>" +
				"		<div class='download' onclick='Download(\"" + FilePath + "/" + f.Name() + "\")'>下载</div>" +
				"		<div class='delete' onclick='Delete(\"" + FilePath + "/" + f.Name() + "\")'>删除</div>" +
				"	</div>"
		}
	}
	return names
}

//组装并返回数据
func templateHtml1(Files string, Name string, response http.ResponseWriter) {
	html := template.New("Index")
	html.Parse(index)
	data := map[string]string{
		"Name": Name,
		"Body": Files,
	}
	html.Execute(response, data)
}
func templateHtml2(Name string, response http.ResponseWriter) {
	html := template.New("Welcome")
	html.Parse(welcome)
	data := map[string]string{
		"Name": Name,
	}
	html.Execute(response, data)
}
