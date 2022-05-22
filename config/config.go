package config

import (
	"log"

	"gopkg.in/ini.v1"
)

// 解析配置文件
var (
	AppMode    string // 服务器启动模式默认 debug 模式
	Port       string //服务启动端口
	JwtKey     string //JWT 签名
	Dbtype     string //数据库类型
	DbHost     string //数据库服务器主机
	DbPort     string //数据服务器端口
	DbUser     string //数据库用户
	DbPassWord string //数据库密码

	DbName string //数据库名

)

func init() {
	f, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatal("配置文件初始化失败")
	}

	loadServer(f)
	loadDb(f)

}

// loadServer 加载服务器配置
func loadServer(file *ini.File) {
	s := file.Section("server")
	AppMode = s.Key("AppMode").MustString("debug")
	Port = s.Key("Port").MustString("3001")
	JwtKey = s.Key("JwtKey").MustString("DouYin")

}

// loadDb 加载数据库相关配置
func loadDb(file *ini.File) {
	s := file.Section("database")
	Dbtype = s.Key("Dbtype").MustString("mysql")
	DbName = s.Key("DbName").MustString("test01")
	DbPort = s.Key("DbPort").MustString("DbPort")
	DbHost = s.Key("DbHost").MustString("DbHost")
	DbUser = s.Key("DbUser").MustString("root")
	DbPassWord = s.Key("DbPassWord").MustString("root")
}

