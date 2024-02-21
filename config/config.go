package config

import (
	"github.com/spf13/viper"
)

func InitConfigs() (*Config, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	Adapter  `mapstructure:"adapter"`
	Logger   `mapstructure:"logger"`
}

type Server struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	WriteTimeOut int64  `mapstructure:"write-timeout"`
	ReadTimeOut  int64  `mapstructure:"read-timeout"`
}

type Database struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	SSLMode   string `mapstructure:"ssl-mode"`
	Migration bool   `mapstructure:"migration"`
}

type Adapter struct {
	//Url string `mapstructure:"url"`
	//Alias string `mapstructure:"alias"`
	//KeyAddress string `mapstructure:"key-address"`
	//PassForKey string `mapstructure:"pass-for-key"`
	//CrtAddress string `mapstructure:"crt-address"`
	//ServiceName string `mapstructure:"service-name"`
	//Timeout int64 `mapstructure:"timeout"`
	//PrivateKeyID string `mapstructure:"private-key-id"`
	//PublicKeyID string `mapstructure:"public-key"`
}

type Logger struct {
	WriteToFile bool   `mapstructure:"write-to-file"`
	Format      string `mapstructure:"format"`
}
