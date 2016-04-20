package main

import (
	"caesarproxy/proxy"
	"caesarproxy/config"
	log "github.com/Sirupsen/logrus"
	"flag"
	"runtime"
	"fmt"
	"os"
)

var (
 	cfg *string = flag.String("c", "/home/jinlong/project/src/caesarproxy/config.yml", "config file")
	vfg *bool = flag.Bool("v", false, "version")
)

const (
	version   = "0.1"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	if *vfg {
		fmt.Println(version)
		os.Exit(0)
	}
	log.Infoln("run caesar-proxy...")
	config.Init(*cfg)
	proxy.Run()
}
