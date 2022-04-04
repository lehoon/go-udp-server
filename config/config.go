package config

import (
	"errors"
	"github.com/spf13/viper"
)

var (
	config *AgenctConfig
)

type Daemon struct {
	Pid  string   `yaml:"pid"`
	Log  string   `yaml:"log"`
}

//tcp server
type UdpServer struct {
	Addr string `yaml:"addr"`
	Port int16  `yaml:"port"`
}

type Local struct {
	Addr    string   `yaml:"addr"`
	UdpServer string `yaml:"udpServer"`
}

type Logger struct {
	MaxAge       int      `yaml:"maxAge"`
	MaxSize      int      `yaml:"maxSize"`
	MaxBackup    int      `yaml:"maxBackup"`
	Compress     bool     `yaml:"compress"`
	Level        int      `yaml:"level"`
	LogPath      string   `yaml:"logPath"`
	ServiceName  string   `yaml:"serviceName"`
}

type AgenctConfig struct {
	Daemon
	Local
	Logger
	UdpServer
}

func init()  {
	NewConfig()
	LoadConfig()
}

func NewConfig() *AgenctConfig {
	config = &AgenctConfig{}
	return config
}

func GetTcpServerPort() int16 {
	return config.UdpServer.Port
}

func GetLocalAddr() string {
	return config.Local.Addr
}

func GetUdpServer() string {
	return config.Local.UdpServer
}

func GetLoggerLevel() int  {
	return config.Logger.Level
}
func GetLoggerPath() string {
	return config.Logger.LogPath
}

func GetLoggerServiceName() string {
	return config.Logger.ServiceName
}

func GetLoggerMaxAge() int {
	return config.Logger.MaxAge
}

func GetLoggerMaxSize() int {
	return config.Logger.MaxSize
}

func GetLoggerMaxBackup() int {
	return config.Logger.MaxBackup
}

func GetLoggerCompress() bool {
	return config.Logger.Compress
}

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("server.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(errors.New("反序列化配置文件出错"))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(errors.New("反序列化配置文件出错"))
	}
}
