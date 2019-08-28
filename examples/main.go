package main

import (
	"github.com/ryannguyen89/go-config"
	"log"
)

func main() {
	config.SetPath("../conf/")
	config.Debug()
	log.Printf(config.Str("FOO"))
	log.Println(config.Str("OBJECT_FIELD1"))
}
