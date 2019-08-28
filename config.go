package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
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
	value := viper.GetString(key)
	if len(value) == 0 {
		return h.getStringWithComplicatedKey(key)
	}
	return value
}

func (h *hConfig) getStringWithComplicatedKey(key string) string {
	subKeys := strings.Split(key, "_")
	for idx, subKey := range subKeys {
		if idx < (len(subKeys) - 1) {
			values := viper.Get(subKey)
			valueMap := values.(map[string]interface{})
			v, ok := valueMap[strings.ToLower(subKeys[idx + 1])]
			if ok {
				return v.(string)
			}
		}
	}

	return ""
}

func (h *hConfig) Get(key string) interface{} {
	return viper.Get(key)
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

func Get(key string) interface{} {
	return getInstance().Get(key)
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
