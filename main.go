package main

import (
	"flag"
	"fmt"
	"go-wallet/api"
	"go-wallet/config"
	"go-wallet/db"
	"go-wallet/util/log"

	"time"
)

var (
	debug    bool
	debugSQL bool
)

func init() {
	fmt.Print("init in main.go\n")
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.BoolVar(&debugSQL, "debugsql", false, "enable sql debug mode")
	config.Load(false, debug, debugSQL)
	log.Init(config.IsDebugMode())
	db.Init()

	time.Sleep(2 * time.Second)
}

func main() {
	api.Start()
}
