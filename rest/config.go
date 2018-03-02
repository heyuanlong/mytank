package rest

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	COOKIE_AUTH_KEY = "_ak"
	TABLE_PREFIX = "tank10_"
	VERSION = "1.0.2"
)

var (
 	CONFIG = &Config{
 		ServerPort :6010,
 		LogToConsole:true,
 		LogPath:"",
 		MatterPath: "",

		MysqlPort: 3306,
		MysqlHost: "127.0.0.1",
		MysqlSchema: "tank",
		MysqlUsername: "tank",
		MysqlPassword: "tank123",
	
		MysqlUrl: "%MysqlUsername:%MysqlPassword@tcp(%MysqlHost:%MysqlPort)/%MysqlSchema?charset=utf8&parseTime=True&loc=Local",

		AdminUsername: "admin",
		AdminEmail: "admin@tank.eyeblue.cn",
		AdminPassword: "123456",
 	}
) 

type Config struct{
	ServerPort int
	LogToConsole bool
	LogPath string
	MatterPath string

	MysqlPort int
	MysqlHost string
	MysqlSchema string
	MysqlUsername string
	MysqlPassword string
	MysqlUrl string

	AdminUsername string
	AdminEmail string
	AdminPassword string
}

func (this *Config) validate() {
	if this.ServerPort == 0{
		LogPanic("ServerPort 未配置")
	}
	if this.MysqlUsername == "" {
		LogPanic("MysqlUsername 未配置")
	}
	if this.MysqlPassword == "" {
		LogPanic("MysqlPassword 未配置")
	}
	if this.MysqlHost == "" {
		LogPanic("MysqlHost 未配置")
	}
	if this.MysqlPort == 0 {
		LogPanic("MysqlPort 未配置")
	}
	if this.MysqlSchema == "" {
		LogPanic("MysqlSchema 未配置")
	}
	this.MysqlUrl = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", this.MysqlUsername, this.MysqlPassword, this.MysqlHost, this.MysqlPort, this.MysqlSchema)
}

//第一级. 从配置文件conf/tank.json中读取配置项
func LoadConfigFromFile() {
	filePath := GetConfPath() + "/tank.json"
	content,err := ioutil.ReadFile(filePath)
	if err != nil {
		LogWarning(fmt.Sprintf("无法找到配置文件：%s,错误：%v\n将使用config.go中的默认配置项。", filePath, err))
	}else{
		err := json.Unmarshal(content,CONFIG)
		if err != nil {
			LogPanic("配置文件格式错误！")
		}

	}
}

func PrepareConfigs() {
	LoadConfigFromFile()

	if CONFIG.LogPath == "" {
		CONFIG.LogPath = GetHomePath() + "/log"
	}
	MakeDirAll(CONFIG.LogPath)

	if CONFIG.MatterPath == "" {
		CONFIG.MatterPath = GetHomePath() + "/matter"
	}
	MakeDirAll(CONFIG.MatterPath)

	//验证配置项的正确性
	CONFIG.validate()
}