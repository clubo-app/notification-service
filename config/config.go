package config

import "github.com/spf13/viper"

type Config struct {
	PORT          int    `mapstructure:"PORT"`
	NATS_CLUSTER  string `mapstructure:"NATS_CLUSTER"`
	SMTP_PW       string `mapstructure:"SMTP_PW"`
	SMTP_USERNAME string `mapstructure:"SMTP_USERNAME"`
	SMTP_HOST     string `mapstructure:"SMTP_HOST"`
	SMTP_PORT     int    `mapstructure:"SMTP_PORT"`
	EMAIL_FROM    string `mapstructure:"EMAIL_FROM"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
