package Upload

import (
	"FileShare/src/AppSet"
	"FileShare/src/SpaceDate"
	"fmt"
	"io"
	"net/http"
	"os"
)

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
		f, err := os.Create(AppSet.GetData() + "/" + request.Form["dir"][0] + "/" + handler.Filename)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintf(response, "done")
		//更新时间
		SpaceDate.UpDate(request.Form["dir"][0])
	} else {
		fmt.Fprintf(response, "非法的访问")
	}
}
