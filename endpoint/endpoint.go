//Package endpoint define and implements endpoint.
package endpoint

import (
	"caesarproxy/config"
)

var (
	EPoints map[string]EndPoint
)

type EndPoint struct {
	Host   string //e.g. http://127.0.0.1 https://google.com
	Port   string
	Weight string
	Health int //health condition: 0-good 1-host can't connect 2-web application can't connect 3-high load
}

func (ep *EndPoint) CheckHealth(mode string) (cond int, err error) {
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
	case "":
		return checkHealthByHost()
	case "":
		return checkHealthByPort()
	case "":
		return checkHealthByURL()
	default:
		return checkHealthByHost()
	}
}

func NewEndPoint(host string, port string, weight string) (ep EndPoint) {
	ep = EndPoint{
		Host:   host,
		Port:   port,
		Weight: weight,
		Health: HealthGood,
	}
	return
}

func init() {
	for _, v := range config.C.Endpoints {
		endpoint := NewEndPoint(v.Host, v.Port, v.Weight)
		EPoints[v.Host] = endpoint
	}
}
