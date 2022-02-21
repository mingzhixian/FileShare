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
	if prefix == "txt" || prefix == "css" || prefix == "js" || prefix == "html" || prefix == "py" || prefix == "java" || prefix == "go" || prefix == "c" || prefix == "cpp" || prefix == "cs" {
		attachment = "<link href=\"./Static/css/txt.css\" type=\"text/css\" rel=\"stylesheet\">"
		return viewTxt(File)
	} else if prefix == "md" {
		attachment = "<link href=\"./Static/css/md.css\" type=\"text/css\" rel=\"stylesheet\">"
		return viewMarkdown(File)
	} else if prefix == "mp4" {
		attachment = "<link href=\"./Static/css/video.css\" type=\"text/css\" rel=\"stylesheet\">"
		return viewVideo(File)
	} else if prefix == "mp3" {
		attachment = "<link href=\"./Static/css/audio.css\" type=\"text/css\" rel=\"stylesheet\">"
		return viewAudio(File)
	} else if prefix == "audio" {
		attachment = "<link href=\"./Static/css/txt.css\" type=\"text/css\" rel=\"stylesheet\">"
		return viewMarkdown(File)
	} else if prefix == "jpg" || prefix == "png" || prefix == "svg" || prefix == "webp" {
		attachment = "<link href=\"./Static/css/img.css\" type=\"text/css\" rel=\"stylesheet\">"
		return viewImg(File)
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
	return "<video id=\"video\" controls src=\"./Download?file=" + File + "\"></video>"
}

func viewAudio(File string) string {
	return "<audio id=\"audio\" src=\"./Download?file=" + File + "\"></audio>"
}

func viewImg(File string) string {
	return "<img id=\"img\" src=\"./Download?file=" + File + "\"></img>"
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
