package config

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var C = Config{}

type Config struct {
	Proxy struct {
		Port    string `yaml:"port"`
		Strategy string `yaml:"strategy"`
		Timeout struct {
			Read  int `yaml:"read"`
			Write int `yaml:"write"`
		} `yaml:"proxy"`
	} `yaml:"proxy"`
	Health struct {
		Type string `yaml:"type"`
		URL string `yaml:"url"`
	       }`yaml:"health"`
	Endpoints []struct{
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Weight string `yaml:"weight"`
	}`yaml:"endpoints"`
}

func ParseConfig(path string) (err error) {
	var b []byte
	var f *os.File
	f, err = os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		//log.Errorln("open config file error:", err)
		return
	}
	b, err = ioutil.ReadAll(f)
	if err != nil {
		//log.Errorln("read config file error:",err)
		return
	}
	//log.Infoln(string(b))
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
	log.Infoln("parse config file. path:", path)
	err := ParseConfig(path)
	if err != nil {
		panic(err)
	}
	log.Infoln("parse config success.")
	log.Debugf("config: \n%v\n ", C)
}
