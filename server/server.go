package server

import (
	"net/http"
	"time"

	"github.com/switfs/oms/config"
	"github.com/switfs/oms/routers"
)

func RunNewServer() {
	Router := routers.InitRouter()

	srv := &http.Server{
		Addr:         config.SR.Server.Hosts + ":" + config.SR.Server.Ports,
		Handler:      Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	switch config.SR.Server.EnabledSSL {
	case true:
		srv.ListenAndServeTLS(config.SR.Server.CertFile, config.SR.Server.CertFile)
	case false:
		srv.ListenAndServe()
	default:
		srv.ListenAndServe()
	}

}
