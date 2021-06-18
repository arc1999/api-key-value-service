package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// AppConfig is ...
type AppConfig struct {
}

// LoadEnv configuration from env file and set it to system environment variable.
func (config AppConfig) LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorln(err)
	}
	keys := viper.AllKeys()
	//log.Println(viper.AllSettings())
	for _, k := range keys {
		key := strings.ToUpper(k)
		envValue, found := os.LookupEnv(key)
		if !found || envValue == "" {
			os.Setenv(key, viper.GetString(k))
		}
	}
}
