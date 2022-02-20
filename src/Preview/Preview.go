package Preview

import (
	"FileShare/src/AppSet"
	_ "embed"

	blackfriday "github.com/russross/blackfriday/v2"

	"io/ioutil"
	"net/http"
	"path"
	"text/template"
	"time"
)

//go:embed preview.html
var preview string

//附加样式文件
var attachment = ""

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

//检测文件类型，并解析
func viewfile(File string) string {
	prefix := path.Ext(File)[1:]
	attachment = "<link href=\"./Static/css/" + prefix + ".css\" type=\"text/css\" rel=\"stylesheet\">"
	if prefix == "txt" {
		return viewTxt(File)
	} else if prefix == "md" {
		return viewMarkdown(File)
	} else if prefix == "mp4" || prefix == "webp" {
		return viewMarkdown(File)
	} else if prefix == "video" {
		return view(File)
	} else if prefix == "audio" {
		return viewMarkdown(File)
	} else if prefix == "img" {
		return viewMarkdown(File)
	} else {
		return "暂不支持此格式预览"
	}
}

//预览函数
func viewTxt(File string) string {
	f, err := ioutil.ReadFile(AppSet.GetData() + "/" + File)
	if err != nil {
		return "发生错误，读取文件失败"
	}
	return "<pre>" + string(f) + "</pre>"
}

func viewMarkdown(File string) string {
	f, err := ioutil.ReadFile(AppSet.GetData() + "/" + File)
	if err != nil {
		return "发生错误，读取文件失败"
	}
	return "<div class=\"markdown-body\">" + string(blackfriday.Run(f)) + "</div>"
}

func viewVideo(File string) string {
	return "<video id=\"video\" src=\"./Download?file=" + File + "\"></video>"
}

//组装并返回数据
func templateHtml(File string, Name string, Body string, response http.ResponseWriter) {
	html := template.New("Preview")
	html.Parse(preview)
	data := map[string]string{
		"Name": Name,
		"File": File,
		"Body": Body,
		"css":  css(),
	}
	html.Execute(response, data)
}

//设置网页主题
func css() string {
	hour := time.Now().Hour()
	if hour > 18 || hour < 8 {
		return "<link href=\"./Static/css/night.css\" type=\"text/css\" rel=\"stylesheet\">" + attachment
	} else {
		return "<link href=\"./Static/css/day.css\" type=\"text/css\" rel=\"stylesheet\">" + attachment
	}
}
