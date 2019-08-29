package main

import (
	"github.com/ryannguyen89/go-config"
	"log"
)

func main() {
	config.SetPath("../conf/")
	config.SetRunModeEnv("MODE")
	config.Debug()
	log.Printf(config.Str("FOO"))
	log.Println(config.Str("OBJECT_FIELD1"))
	log.Println(config.GetRunMode())
}
