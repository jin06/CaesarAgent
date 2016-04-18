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
	vflag *bool = flag.Bool("v", false, "version")
	version = "0.1"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	if *vflag {
		fmt.Println(version)
		os.Exit(0)
	}
	log.Infoln("caesar-proxy start...")
	config.Init(*cfg)
	proxy.Run()
}
