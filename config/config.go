package config

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

var C = Config{}

type Config struct {
	Proxy struct {
		Port    string `yaml:"port"`
		Strategy string `yaml:"strategy"`
		Timeout struct {
			Read  time.Duration `yaml:"read"`
			Write time.Duration `yaml:"write"`
		} `yaml:"timeout"`
	} `yaml:"proxy"`
	Health struct {
		Type string `yaml:"type"`
		URL string `yaml:"url"`
	       }`yaml:"health"`
	Servers []struct{
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Weight string `yaml:"weight"`
	}`yaml:"servers"`
	Environment string `yaml:"environment"`
}

func ParseConfig(path string) (err error) {
	var b []byte
	var f *os.File
	f, err = os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	b, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &C)
	return
}

func Init(path string) {
	defer func() {
		if e := recover(); e!=nil {
			log.Errorln("parse config file error:", e)
			os.Exit(0)
		}
	}()
	log.Infoln("Parse config file, path:", path)
	err := ParseConfig(path)
	if err != nil {
		panic(err)
	}
	log.Infoln("parse config success.")
	log.Infof("config: \n%+v\n ", C)
}
