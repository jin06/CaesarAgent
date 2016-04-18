//Package strategy implements load balance strategy.
package strategy

import (
	"net/http"
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
	return
}