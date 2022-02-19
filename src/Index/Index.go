package Index

import (
	"FileShare/src/AppSet"
	_ "embed"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

//go:embed index.html
var index string

//获取文章
func Index(response http.ResponseWriter, request *http.Request) {
	//获取数据文件夹目录和分享站名字
	Files := scanFiles(AppSet.GetData())
	Name := AppSet.GetName()
	//返回数据
	templateHtml(Files, Name, response)
}

//扫描文件夹下的所有文件，返回文件名的集合
func scanFiles(FilePath string) string {
	var names string
	files, err := ioutil.ReadDir(FilePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		names += "<div class=\"file\" onclick=\"Preview('" + f.Name() + "')\">" + f.Name() + "</div>\n"
	}
	return names
}

//组装并返回数据
func templateHtml(Files string, Name string, response http.ResponseWriter) {
	html := template.New("Index")
	html.Parse(index)
	data := map[string]string{
		"Name":        Name,
		"Body":        Files,
		"DayAndNight": dayAndNight(),
	}
	html.Execute(response, data)
}

//设置网页主题
func dayAndNight() string {
	hour := time.Now().Hour()
	if hour > 18 || hour < 8 {
		return "./Static/css/night.css"
	} else {
		return "./Static/css/day.css"
	}
}
