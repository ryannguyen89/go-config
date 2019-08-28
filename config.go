package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	defaultConfigFile = "config.yaml"
	defaultPath = "conf/"
)

func init() {
	var (
		runMode = os.Getenv("RUNMODE")
		configFile string
	)

	log.Printf("RUNMODE is: %s\n", runMode)
	if len(runMode) > 0 {
		configFile = fmt.Sprintf("%s.%s", runMode, defaultConfigFile)
	} else {
		configFile = defaultConfigFile
	}
	configFile = defaultPath + configFile

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Read config file error: %v\n", err)
	}
	viper.AutomaticEnv()
}

func Str(key string) string {
	return viper.GetString(key)
}

func Int(key string) int {
	return viper.GetInt(key)
}

func Int32(key string) int32 {
	return viper.GetInt32(key)
}

func Debug() {
	viper.Debug()
}
