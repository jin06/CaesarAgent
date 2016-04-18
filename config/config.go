package config

import (
	"gopkg.in/yaml.v2"
	log "github.com/Sirupsen/logrus"
	"os"
)

var C *Config

type Config struct {
	Proxy struct {
		Port    string `yaml:"port"`
		Timeout struct {
			Read  int `yaml:"read"`
			Write int `yaml:"write"`
		}`yaml:"proxy"`
	}`yaml:"proxy"`
}

func ParseConfig(path string) (err error) {
	var f *os.File
	if f, err = os.Open(path); err !=nil {
		return
	}
	defer f.Close()
	var b []byte
	if _, err = f.Read(b); err !=nil {
		return
	}
	log.Infoln(string(b))
	C = new(Config)
	err = yaml.Unmarshal(b, C)
	return
}

func Init(path string) {
	if err := ParseConfig(path); err !=nil {
		log.Errorln("parse config error:", err)
		os.Exit(0)
	}
}
