package lib

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type BaseConf struct {
	DebugMode    string    `mapstructure:"debug_mode"`
	TimeLocation string    `mapstructure:"time_location"`
	Base         struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
	} `mapstructure:"base"`
}

var ConfBase *BaseConf
var ViperConfMap map[string]*viper.Viper

func InStringArray(s string,arr []string) bool{
	for _,v := range arr{
		if s==v{
			return true
		}
	}
	return false
}

func InitModule(configPath string,modules []string) error {
	var conf *string
	if len(configPath) > 0 {
		conf = &configPath
	} else {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " Can't not find init config:初始化配置为空")
		os.Exit(1)
	}

	// 解析配置文件目录
	if err := ParseConfPath(*conf); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " ParseConfPath:"+err.Error())
	}

	//初始化配置文件
	if err := InitViperConf(); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitViperConf:"+err.Error())
	}

	if InStringArray("base",modules){
		if err := InitBaseConf(GetConfPath(configPath,"base")); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitBaseConf:"+err.Error())
		}
	}

	if InStringArray("mysql",modules){
		if err := InitDBPool(GetConfPath(configPath,"mysql")); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitDBPool:"+err.Error())
		}
	}

	if InStringArray("redis",modules){
		if err := InitRedisConf(GetConfPath(configPath,"redis")); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitRedis:"+err.Error())
		}
	}

	return nil
}

func GetBaseConf() *BaseConf {
	return ConfBase
}

func InitBaseConf(path string) error {
	ConfBase = &BaseConf{}
	err := ParseConfig(path, ConfBase)
	if err != nil {
		return err
	}

	if ConfBase.DebugMode == "" {
		if ConfBase.Base.DebugMode!=""{
			ConfBase.DebugMode = ConfBase.Base.DebugMode
		}else{
			ConfBase.DebugMode = "debug"
		}
	}
	if ConfBase.TimeLocation == "" {
		if ConfBase.Base.TimeLocation!=""{
			ConfBase.TimeLocation = ConfBase.Base.TimeLocation
		}else{
			ConfBase.TimeLocation = "Asia/Guangzhou"
		}
	}

	fmt.Println("ConfBase=",ConfBase)
	return nil
}

func ParseConfPath(config string) error {
	path := strings.Split(config, "/")
	prefix := strings.Join(path[:len(path)-1], "/")
	ConfEnvPath = prefix
	ConfEnv = path[len(path)-2]
	return nil
}

func GetConfFilePath(fileName string) string {
	return ConfEnvPath + "/" + fileName
}

func InitViperConf() error {
	f, err := os.Open(ConfEnvPath + "/")
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, err := ioutil.ReadFile(ConfEnvPath + "/" + f0.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			v.SetConfigType("toml")
			v.ReadConfig(bytes.NewBuffer(bts))
			pathArr := strings.Split(f0.Name(), ".")
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper)
			}
			ViperConfMap[pathArr[0]] = v
		}
	}
	return nil
}

func GetIntConf(key string) int {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetInt(strings.Join(keys[1:len(keys)], "."))
	return conf
}

func GetStringConf(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return ""
	}
	v, ok := ViperConfMap[keys[0]]
	if !ok {
		return ""
	}
	confString := v.GetString(strings.Join(keys[1:len(keys)], "."))
	return confString
}