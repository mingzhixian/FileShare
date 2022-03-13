package Download

import (
	"FileShare/src/AppSet"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Download(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	//获取数据文件夹目录和分享站名字
	filePath := request.Form["dir"][0]
	f, err := ioutil.ReadFile(AppSet.GetData() + "/" + filePath)
	if err != nil {
		fmt.Println(err)
	}
	response.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filePath))
	response.Header().Add("Content-Type", "application/file")

	content := bytes.NewReader(f)
	http.ServeContent(response, request, filePath, time.Now(), content)
}
