//Package endpoint define and implements endpoint.
package server

import (
	"caesarproxy/config"
)

var (
	EPoints map[string]Server
)

type Server struct {
	Host   string //e.g. http://127.0.0.1 https://google.com
	Port   string
	Weight string
	Health int //health condition: 0-good 1-host can't connect 2-web application can't connect 3-high load
}

func (ep *Server) CheckHealth(mode string) (cond int, err error) {
	checkHealthByHost := func() (cond int, err error) {
		return
	}
	checkHealthByPort := func() (cond int, err error) {
		return
	}
	checkHealthByURL := func() (cond int, err error) {
		return
	}
	switch mode {
	case "host":
		return checkHealthByHost()
	case "port":
		return checkHealthByPort()
	case "url":
		return checkHealthByURL()
	default:
		return checkHealthByHost()
	}
}

func NewServer(host string, port string, weight string) (ep Server) {
	ep = Server{
		Host:   host,
		Port:   port,
		Weight: weight,
		Health: HealthGood,
	}
	return
}

func init() {
	for _, v := range config.C.Servers {
		endpoint := NewServer(v.Host, v.Port, v.Weight)
		EPoints[v.Host] = endpoint
	}
}
