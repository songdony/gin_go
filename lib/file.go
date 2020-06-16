package lib

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var ConfEnvPath string //配置文件夹
var ConfEnv string     //配置环境名 比如：dev prod test

func ParseConfig(path string, conf interface{}) error {
	file, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Open config %v fail, %v", path, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Read config fail, %v", err)
	}
	v:=viper.New()
	v.SetConfigType("toml")
	v.ReadConfig(bytes.NewBuffer(data))

	if err:=v.Unmarshal(conf);err!=nil{
		return fmt.Errorf("Parse config fail, config:%v, err:%v", string(data), err)
	}

	return nil
}

func GetConfPath(fileName string) string {
	return  "./" + fileName + ".toml"
}

func GetConfEnv() string{
	return ConfEnv
}