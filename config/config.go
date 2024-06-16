package config

import (
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var Conf *Config

// 定义供整个项目使用的变量
const ()

type Config struct {
	Common Common
	//Connect ConnectConfig
	Logic LogicConfig
	//Task    TaskConfig
	//Api     ApiConfig
	//Site    SiteConfig
}

var once sync.Once

func init() {
	initConfig()
}

// use viper to extract config
func initConfig() {
	once.Do(func() {
		realPath := getCurrentDir()
		configFilePath := realPath + "/" + "dev" + "/"
		viper.SetConfigType("toml")
		viper.AddConfigPath(configFilePath)
		viper.SetConfigName("logic")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		Conf = new(Config)
		viper.Unmarshal(&Conf.Logic)
	})
}
func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(1)
	aPath := strings.Split(fileName, "/")
	dir := strings.Join(aPath[0:len(aPath)-1], "/")
	return dir
}

// logic
type LogicConfig struct {
	LogicBase LogicBase `mapstructure:"logic-base"`
}

type LogicBase struct {
	ServerId   string `mapstructure:"serverId"`
	CpuNum     int    `mapstructure:"cpuNum"`
	RpcAddress string `mapstructure:"rpcAddress"`
	CertPath   string `mapstructure:"certPath"`
	KeyPath    string `mapstructure:"keyPath"`
}

// common
type CommonEtcd struct {
	Host              string `mapstructure:"host"`
	BasePath          string `mapstructure:"basePath"`
	ServerPathLogic   string `mapstructure:"serverPathLogic"`
	ServerPathConnect string `mapstructure:"serverPathConnect"`
	UserName          string `mapstructure:"userName"`
	Password          string `mapstructure:"password"`
	ConnectionTimeout int    `mapstructure:"connectionTimeout"`
}
type CommonRedis struct {
	RedisAddress  string `mapstructure:"redisAddress"`
	RedisPassword string `mapstructure:"redisPassword"`
	Db            int    `mapstructure:"db"`
}
type Common struct {
	CommonEtcd  CommonEtcd  `mapstructure:"common-etcd"`
	CommonRedis CommonRedis `mapstructure:"common-redis"`
}
