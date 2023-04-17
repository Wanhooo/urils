package config

import (
	"encoding/json"
	"io/ioutil"
)

// 外界go开发项目： 配置文件都是 json、yaml、ini

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}


// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver       string `json:"driver"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Database     string `json:"database"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Charset      string `json:"charset"`
	MaximumConn     int `json:"maximum_connection"`
	MaximumFreeConn int `json:"maximum_free_connection"`
	TimeOut         int `json:"timeout"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Database    int    `json:"database"`
	Password    string `json:"password"`
	MaxConn     int `json:"max_connection"`
	MaxFreeConn int `json:"max_free_connection"`
	TimeOut     int `json:"timeout"`
}

type RedisConfigGroup struct {
	Default    RedisConfig
	Sms        RedisConfig
}

// Config 整个项目的配置
type Config struct {
	Mode       string `json:"mode"`
	Port       int    `json:"port"`
	*LogConfig `json:"log"`
	*DatabaseConfig `json:"database"`
	*RedisConfigGroup `json:"redis"`
}

// Conf 全局配置变量
var Conf = new(Config)


// Init 初始化配置；从指定文件加载配置文件
func Init(filePath string) error {
	/**
	 filePath 配置文件json文件的路径
	 */
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf)
}
