package Delete

import (
	"fmt"
	"net/http"
	"os"
)

func Delete(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	err := os.Remove(request.Form["filePath"][0])
	if err != nil {
		fmt.Println(err)
	}
}
