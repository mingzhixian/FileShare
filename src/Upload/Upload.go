package Upload

import (
	"FileShare/src/AppSet"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

//go:embed upload.html
var upload string

func Upload(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	if request.Method == "POST" {
		//获取文件
		file, handler, err := request.FormFile("file")
		if err != nil {
			http.Error(response, err.Error(), 500)
			return
		}
		defer file.Close()
		//写入文件
		f, err := os.Create(AppSet.GetData() + "/" + handler.Filename)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintf(response, "done")
	} else if request.Method == "GET" && len(request.Form["NewFiles"]) != 0 {
		err := os.MkdirAll(AppSet.GetData()+"/"+request.Form["NewFiles"][0], os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(response, "done")
	} else {
		Name := AppSet.GetName()
		html := template.New("Preview")
		html.Parse(upload)
		data := map[string]string{
			"Name":        Name,
			"DayAndNight": dayAndNight(),
		}
		html.Execute(response, data)
	}
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
