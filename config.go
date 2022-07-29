package main

import (
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type ConfigYaml struct {
	Width  string `mapstructure:"width"`
	Height string `mapstructure:"height"`
}

type Config struct {
	Width  int
	Height int
}

var vp *viper.Viper

func loadConfig(args ...string) Config {
	var confYaml ConfigYaml

	vp = viper.New()
	vp.SetConfigFile("config.yml")
	vp.AddConfigPath(".")
	for _, el := range args {
		vp.AddConfigPath(el)
	}
	err := vp.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read the config file: ", err.Error())
		return Config{}
	}

	err = vp.Unmarshal(&confYaml)
	if err != nil {
		log.Fatal("Cannot unmarshal the config file: ", err.Error())
		return Config{}
	}

	w, err := strconv.Atoi(confYaml.Width)
	h, err := strconv.Atoi(confYaml.Height)
	return Config{Width: w, Height: h}
}
