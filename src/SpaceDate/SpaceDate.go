package SpaceDate

import (
	"FileShare/src/AppSet"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func UpDate(filePath string) {
	site := strings.IndexAny(filePath, "/")
	usrSpace := ""
	if site != -1 {
		usrSpace = filePath[:site]
	} else {
		usrSpace = filePath
	}
	f, err := os.OpenFile(AppSet.GetData()+"/"+usrSpace+"/空间守则", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	write(f, "10")
	defer f.Close()
}
func DownDate() {
	files, err := ioutil.ReadDir(AppSet.GetData())
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		if f.IsDir() {
			file, err := os.OpenFile(AppSet.GetData()+"/"+f.Name()+"/空间守则", os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			day := getDay(file) - 1
			file, err = os.OpenFile(AppSet.GetData()+"/"+f.Name()+"/空间守则", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				fmt.Println(err)
			}
			if day == 0 {
				err := os.RemoveAll(AppSet.GetData() + "/" + f.Name())
				if err != nil {
					fmt.Println(err)
				}
			} else {
				write(file, strconv.Itoa(day))
			}
		}
	}
}
func getDay(f *os.File) int {
	reader := bufio.NewReader(f)
	day, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		day = day[strings.IndexAny(day, "：")+3 : strings.IndexAny(day, "\n")]
	}
	intday, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println(err)
	}
	return intday
}
func write(f *os.File, day string) {
	write := bufio.NewWriter(f)
	write.WriteString("                剩余生存时间：" + day + "\n" +
		"        |-----------------------------------|\n" +
		"        | 公益空间，请勿实施恶意行为，空间不 |\n" +
		"        | 限制大小，但长时间不操作（不包括下 |\n" +
		"        | 载）会过期，请注意不要操作他人的文 |\n" +
		"        | 件。本文件分享站为公开站，请勿存放 |\n" +
		"        | 重要资料。本站不对文件的安全负责。 |\n" +
		"        |-----------------------------------|\n")
	write.Flush()
}
