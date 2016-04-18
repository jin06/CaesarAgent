package proxy

import (
	"net/http"
	"time"
	//log "github.com/Sirupsen/logrus"
	sty "caesarproxy/proxy/strategy"
)

type Server struct {
	Server *http.Server
	Strategy sty.Strategy
}

func newServer(port string,sty sty.Strategy) (*Server){
	sm := http.NewServeMux()
	sm.HandleFunc("/",sty.Redirect)
	srv := &http.Server{
		Addr:":"+port,
		Handler:sm,
		ReadTimeout: 10 * time.Second,
		WriteTimeout:10 * time.Second,
		MaxHeaderBytes: 1 << 20, //1M
	}
	server := &Server{
		Server:srv,
		Strategy:sty,
	}
	return server
}