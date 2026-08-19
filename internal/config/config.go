package config

import (
	"tracker/internal/api/http"
	"tracker/internal/base/database"

	"github.com/spf13/viper"
)

type Config struct {
	Server *http.Config
	MySQL  *database.MySQLConfig
	Redis  *database.RedisConfig
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	return &Config{
		Server: &http.Config{
			Host: viper.GetString("SERVER_HOST"),
			Port: viper.GetString("SERVER_PORT"),
		},
		MySQL: &database.MySQLConfig{
			Host:     viper.GetString("MYSQL_HOST"),
			Port:     viper.GetString("MYSQL_PORT"),
			DBName:   viper.GetString("MYSQL_DBNAME"),
			Username: viper.GetString("MYSQL_USERNAME"),
			Password: viper.GetString("MYSQL_PASSWORD"),
		},
		Redis: &database.RedisConfig{
			Addr:     viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
	}, nil
}
