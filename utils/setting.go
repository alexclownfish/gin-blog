package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db       string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPasswd string
	DbName   string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
	//QiniuSoftServer  string
	//QiniuVideoServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3003")
	JwtKey = file.Section("server").Key("JwtKey").MustString("678dsnakg6234")
}
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPasswd = file.Section("database").Key("DbPasswd").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("gin-blog")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
	//QiniuSoftServer = file.Section("qiniu").Key("QiniuSoftServer").String()
	//QiniuVideoServer = file.Section("qiniu").Key("QiniuVideoServer").String()
}
