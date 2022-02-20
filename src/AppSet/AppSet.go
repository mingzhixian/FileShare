package AppSet

import (
	"fmt"
	"os"
)

var Data = ""
var Name = ""
var MaxSize int64 = 1024

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

//设置最大允许文件大小
func SetMaxSize(maxSize int64) {
	MaxSize = maxSize
}

//获取最大允许文件大小
func GetMaxSize() int64 {
	return MaxSize
}
