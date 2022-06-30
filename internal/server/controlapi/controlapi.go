package controlapi

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/k0tletka/SDFS/internal/server/config"
	"github.com/k0tletka/SDFS/internal/server/controlapi/handler"
	"net/http"
	"time"
)

type ControlAPI struct {
	ctx  context.Context
	conf *config.ServerConfig

	httpServer *http.Server
	inited     bool
}

func (s *ControlAPI) Init(ctx context.Context, conf *config.ServerConfig) {
	s.ctx, s.conf = ctx, conf

	apiHandler := handler.NewControlAPIHandler(ctx, conf)

	// Init http server
	s.httpServer = &http.Server{
		Handler: getControlAPIRouter(apiHandler),
		Addr: fmt.Sprintf(
			"%s:%d",
			conf.Config.ControlAPIConf.ListenAddr,
			conf.Config.ControlAPIConf.ListenPort,
		),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.inited = true
}

func (s *ControlAPI) StartServer() error {
	if !s.inited {
		panic("Server must be inited before starting")
	}

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		if s.conf.Config.ControlAPIConf.UseSSL {
			errChan <- s.httpServer.ListenAndServeTLS(
				s.conf.Config.ControlAPIConf.CertFilePath,
				s.conf.Config.ControlAPIConf.KeyFilePath,
			)
		} else {
			errChan <- s.httpServer.ListenAndServe()
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-s.ctx.Done():
		s.httpServer.Shutdown(context.Background())
		return context.Canceled
	}
}

func getControlAPIRouter(handlers *handler.ControlAPIHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")

	return router
}
