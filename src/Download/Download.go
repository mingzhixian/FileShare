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
	File := request.Form["file"][0]
	f, err := ioutil.ReadFile(AppSet.GetData() + "/" + File)
	if err != nil {
		fmt.Println(err)
	}
	response.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, File))
	response.Header().Add("Content-Type", "application/file")

	content := bytes.NewReader(f)
	http.ServeContent(response, request, File, time.Now(), content)
}
