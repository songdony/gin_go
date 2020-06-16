package lib

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//var db *gorm.DB

//var DBMapPool map[string]*gorm.DB
var GORMMapPool map[string]*gorm.DB
//var DBDefaultPool *gorm.DB
//var GORMDefaultPool *gorm.DB
//var ViperConfMap map[string]*viper.Viper

type MysqlMapConf struct {
	List map[string]*MySQLConf `mapstructure:"list"`
}

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

func InitDBPool(path string) error{
	DbConfMap := &MysqlMapConf{}
	err := ParseConfig(path, DbConfMap)

	GORMMapPool = map[string]*gorm.DB{}   // 先初始化不然报错 panic: assignment to entry in nil map

	for confName, DbConf := range DbConfMap.List {
		//gorm连接方式
		var dbgorm *gorm.DB
		dbgorm, err := gorm.Open("mysql", DbConf.DataSourceName)
		if err != nil {
			return err
		}
		dbgorm.SingularTable(true)
		dbgorm.LogMode(true)

		err = dbgorm.DB().Ping()
		if err != nil {
			return err
		}

		GORMMapPool[confName] = dbgorm
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func  GetGORMMapPool(name string) (*gorm.DB,error) {
	if dbpool,ok := GORMMapPool[name];ok{
		return dbpool,nil
	}
	return nil,errors.New("get pool error")
}

func GetGormPool(name string) (*gorm.DB, error) {
	if dbpool, ok := GORMMapPool[name]; ok {
		return dbpool, nil
	}
	return nil, errors.New("get pool error")
}

func CloseDB() error {
	for _,dbpool := range GORMMapPool{
		dbpool.Close()
	}
	return nil
}

