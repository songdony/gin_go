package lib

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"time"
)

type RedisMapConf struct {
	List map[string]*RedisConf `mapstructure:"list"`
}

type RedisConf struct {
	ProxyList []string `mapstructure:"proxy_list"`
	MaxActive int      `mapstructure:"max_active"`
	MaxIdle   int      `mapstructure:"max_idle"`
	Downgrade bool     `mapstructure:"down_grade"`
}

var ConfRedis *RedisConf
var ConfRedisMap *RedisMapConf


func InitRedisPool(){
	if err := InitRedisConf(GetConfPath("redis_map")); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitRedisConf:"+err.Error())
	}
}

func InitRedisConf(path string) error {
	ConfRedis := &RedisMapConf{}
	err := ParseConfig(path, ConfRedis)
	if err != nil {
		return err
	}
	ConfRedisMap = ConfRedis
	return nil
}

func RedisConnFactory(name string) (redis.Conn, error) {
	for confName, cfg := range ConfRedisMap.List {
		if name == confName {
			randHost := cfg.ProxyList[rand.Intn(len(cfg.ProxyList))]
			return redis.Dial(
				"tcp",
				randHost,
				redis.DialConnectTimeout(50*time.Millisecond),
				redis.DialReadTimeout(100*time.Millisecond),
				redis.DialWriteTimeout(100*time.Millisecond))
		}
	}
	return nil, errors.New("create redis conn fail")
}