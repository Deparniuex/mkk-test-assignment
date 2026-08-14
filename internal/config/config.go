package config

type Config struct {
	HttpHost string
	HttpPort string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
}

type RedisConfig struct {
}

func NewConfig() (*Config, error) {
	return &Config{}, nil
}
