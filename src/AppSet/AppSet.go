package AppSet

import (
	"fmt"
	"os"
)

var Data = ""
var Name = ""

//设置数据地址
func SetData(path string) {
	Data = path
	//创建数据文件夹
	err := os.MkdirAll(Data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

//获取数据地址
func GetData() string {
	return Data
}

//设置分享站名字
func SetName(name string) {
	Name = name
}

//获取分享站名字
func GetName() string {
	return Name
}
