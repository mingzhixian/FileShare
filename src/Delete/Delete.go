package Delete

import (
	"FileShare/src/AppSet"
	"FileShare/src/SpaceDate"
	"fmt"
	"net/http"
	"os"
)

func Delete(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	err := os.RemoveAll(AppSet.GetData() + "/" + request.Form["dir"][0])
	if err != nil {
		fmt.Println(err)
	}
	//更新时间
	SpaceDate.UpDate(request.Form["dir"][0])
}
