package http

import (
	"fmt"
	"net/http"
	"platform/config"
	"platform/logging"
	"platform/pipeline"
	"sync"
)

type PipelineAdaptor struct {
	pipeline.RequestPipeline
}

func (p PipelineAdaptor) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	err := p.ProcessRequest(req, resp)
	if err != nil {
		fmt.Printf("Error while handling request: %v", err)
	}
}

func Serve(pl pipeline.RequestPipeline, cfg config.Configuration, logger logging.Logger) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	adaptor := PipelineAdaptor{RequestPipeline: pl}
	enableHTTP := cfg.GetBoolDefault("http:enableHTTP", true)
	if enableHTTP {
		httpPort := cfg.GetIntDefault("http:port", 5000)
		logger.Debugf("Starting HTTP server on port %v", httpPort)
		wg.Add(1)
		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adaptor)
			if err != nil {
				panic(err)
			}
		}()
	}
	enableHTTPS := cfg.GetBoolDefault("http:enableHTTPS", false)
	if enableHTTPS {
		httpsPort := cfg.GetIntDefault("http:httpsPort", 5500)
		certFile, cfok := cfg.GetString("http:httpsCert")
		keyFile, kfok := cfg.GetString("http:httpsKey")
		if cfok && kfok {
			logger.Debugf("Starting HTTPS server on port %v", httpsPort)
			wg.Add(1)
			go func() {
				err := http.ListenAndServeTLS(fmt.Sprintf(":%v", httpsPort), certFile, keyFile, adaptor)
				if err != nil {
					panic(err)
				}
			}()
		} else {
			panic("HTTPS certificate settings not found")
		}
	}
	return &wg
}
