package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Dest string
	Emoticon map[string]string
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("emote")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = v.Unmarshal(c)
	return c, err
}