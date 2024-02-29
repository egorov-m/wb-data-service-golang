package config

import "github.com/spf13/viper"

type Config struct {
	Env         string      `mapstructure:"ENV"`
	Redis       RedisConfig `mapstructure:"REDIS"`
	Concurrency int         `mapstructure:"CONCURRENCY"`
	DatabaseDsn string      `mapstructure:"DATABASE_DSN"`
}

type RedisConfig struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

func New(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("failed read config: " + err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic("failed unmarshal config: " + err.Error())
	}

	return &config
}
