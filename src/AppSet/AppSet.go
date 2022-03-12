package AppSet

var Data = ""
var Name = ""

//设置数据地址
func SetData(path string) {
	Data = path
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
