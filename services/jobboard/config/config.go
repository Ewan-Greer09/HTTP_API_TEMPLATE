package config

import (
	"log"

	viper "github.com/spf13/viper"
)

type JobBoardConfig struct {
	Port    string
	Address string
}

var Config = &JobBoardConfig{}

func Init() JobBoardConfig {
	viper := viper.New()
	viper.SetConfigName("production")
	viper.AddConfigPath("./services/jobboard/config/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	viper.SetDefault("port", "8080")
	viper.SetDefault("address", "localhost")

	err = viper.Unmarshal(Config)
	if err != nil {
		log.Panic(err)
	}

	// TODO: Config validation
	return *Config
}
