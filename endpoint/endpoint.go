//Package endpoint define and implements endpoint.
package endpoint
import (
	"caesarproxy/config"
)
var (
	EPoints map[string]EndPoint
)

type EndPoint struct {
	Host string  //e.g. http://127.0.0.1 https://google.com
	Port string
	Weight string
	Health int //health condition: 0-good 1-host can't connect 2-web application can't connect 3-high load
}

func NewEndPoint(host string,port string, weight string) EndPoint{
	ep := EndPoint{
		Host:host,
		Port : port,
		Weight : weight,
		Health:HealthGood,
	}
	return ep
}

func init(){
	for _, v := range config.C.Endpoints{
		endpoint := NewEndPoint(v.Host,v.Port,v.Weight)
		EPoints[v.Host] = endpoint
	}
}