package main

import (
	"caesarproxy/config"
	"caesarproxy/proxy"
	_ "caesarproxy/server"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"runtime"
)

var (
	cfg *string = flag.String("config", "./config.yml", "config file")
	vfg *bool   = flag.Bool("version", false, "version")
)

const (
	version = "0.1"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	configLog()

	flag.Parse()

	if *vfg {
		fmt.Println(version)
		os.Exit(0)
	}

	log.Infoln("run caesar-proxy...")

	config.Init(*cfg)

	proxy.Run()
}

func configLog() {
	text := log.TextFormatter{
		ForceColors:      true,
		DisableColors:    false,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableSorting:   false,
	}
	log.SetFormatter(&text)
	log.SetLevel(log.DebugLevel)
}
