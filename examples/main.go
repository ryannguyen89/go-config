package main

import (
	"github.com/hoanghon00/go-config"
	"log"
)

func main() {
	config.SetPath("../conf/")
	config.Debug()
	log.Printf(config.Str("FOO"))
}
