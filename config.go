package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

const (
	defaultConfigFile = "config.yaml"
	defaultPath = "conf/"
	defaultType = "yaml"
)

var (
	configPath string
	configFile string
	configType string
	configInstance *hConfig
	configOnce sync.Once
)

type hConfig struct {}

func (h *hConfig) Str(key string) string {
	return viper.GetString(key)
}

func (h *hConfig) Int(key string) int {
	return viper.GetInt(key)
}

func (h *hConfig) Int32(key string) int32 {
	return viper.GetInt32(key)
}

func (h *hConfig) Debug() {
	viper.Debug()
}

func getInstance() *hConfig {
	configOnce.Do(func() {
		configInstance = initializeConfigInstance()
	})

	return configInstance
}

func initializeConfigInstance() *hConfig {
	var (
		runMode              = os.Getenv("RUNMODE")
		processedConfigFile string
	)
	log.Printf("RUNMODE is: %s\n", runMode)

	if len(configPath) == 0 {
		configPath = defaultPath
	}
	if len(configFile) == 0 {
		configFile = defaultConfigFile
	}
	if len(configType) == 0 {
		configType = defaultType
	}
	if len(runMode) > 0 {
		processedConfigFile = fmt.Sprintf("%s.%s", runMode, configFile)
	} else {
		processedConfigFile = configFile
	}
	processedConfigFile = configPath + processedConfigFile

	viper.SetConfigType(configType)
	viper.SetConfigFile(processedConfigFile)

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Read config file error: %v\n", err)
	}
	viper.AutomaticEnv()

	return &hConfig{}
}

func SetFile(file string) {
	configFile = file
}

func SetPath(path string) {
	configPath = path
}

func SetType(_type string) {
	configType = _type
}

func Str(key string) string {
	return getInstance().Str(key)
}

func Int(key string) int {
	return getInstance().Int(key)
}

func Int32(key string) int32 {
	return getInstance().Int32(key)
}

func Debug() {
	getInstance().Debug()
}
