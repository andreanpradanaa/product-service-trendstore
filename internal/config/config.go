package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Database
}

type Database struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

var cfg Config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	viper.AddConfigPath(basepath)
	viper.SetConfigFile(".yaml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	viper.Unmarshal(&cfg)
}

func GetConfig() *Config {
	return &cfg
}
