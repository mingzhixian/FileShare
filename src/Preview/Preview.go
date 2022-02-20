package Preview

import (
	"FileShare/src/AppSet"
	_ "embed"
	"net/http"
	"text/template"
	"time"
)

//go:embed preview.html
var preview string

func Preview(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	//获取数据文件夹目录和分享站名字
	File := request.Form["file"][0]
	Name := AppSet.GetName()
	Body := viewfile(File)
	if File == "" {
		http.Redirect(response, request, "./", http.StatusFound) //重定向
	} else {
		//返回数据
		templateHtml(File, Name, Body, response)
	}
}

func viewfile(File string) string {
	return "暂不支持此格式预览"
}

//组装并返回数据
func templateHtml(File string, Name string, Body string, response http.ResponseWriter) {
	html := template.New("Preview")
	html.Parse(preview)
	data := map[string]string{
		"Name":        Name,
		"File":        File,
		"Body":        Body,
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
