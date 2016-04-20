//Package endpoint define and implements endpoint.
package endpoint

type EndPoint struct {
	Host string  //e.g. http://127.0.0.1 https://google.com
	Health int //health condition: 0-good 1-host can't connect 2-web application can't connect 3-high load
}

func NewEndPoint(host string) EndPoint{
	ep := EndPoint{
		Host:host,
		Health:HealthGood,
	}
	return ep
}