package config

import (
	"github.com/spf13/viper"
)

//Config recebe os dados do banco de dados
type Config struct {
	Environment string
	Mongo       MongoConfiguration
}

//MongoConfiguration  recebe os dados conexão do mongoDB
type MongoConfiguration struct {
	Server   string
	Database string
}

func (c *Config) Read() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}
