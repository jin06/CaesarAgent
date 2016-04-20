//Package strategy implements load balance strategy.
package strategy

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
)

type Strategy interface{
	Redirect(http.ResponseWriter,*http.Request)
}

type PollStrategy struct {}

func NewPollStrategy() *PollStrategy{
	sty := &PollStrategy{}
	return sty
}

func (ps *PollStrategy) Redirect(w http.ResponseWriter,r *http.Request){
	log.Infoln(r.URL)
	return
}