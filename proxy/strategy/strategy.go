//Package strategy implements load balance strategy.
package strategy

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type Strategy interface {
	Init()
	Redirect(http.ResponseWriter, *http.Request)
}

func redirect(w http.ResponseWriter, host string, port string, redirectTyp int) {
	w.Header().Set("location", host+":"+port)
	w.WriteHeader(redirectTyp)
	log.Infoln("redirect to ", host+":"+port)
}
