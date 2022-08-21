package controlapi

import (
	"context"
	"fmt"
	"github.com/k0tletka/sdfsd/internal/server/config"
	"github.com/k0tletka/sdfsd/internal/server/controlapi/handler"
	"github.com/labstack/echo/v4"
)

type ControlAPI struct {
	ctx  context.Context
	conf *config.ServerConfig

	httpServer *echo.Echo
	inited     bool
}

func (s *ControlAPI) Init(ctx context.Context, conf *config.ServerConfig) {
	s.ctx, s.conf = ctx, conf

	s.httpServer = echo.New()
	s.httpServer.HideBanner = true

	handlers := handler.NewControlAPIHandler(ctx, conf)

	s.registerRoutes(handlers, s.httpServer)
	s.inited = true
}

func (s *ControlAPI) StartServer() error {
	if !s.inited {
		panic("Server must be inited before starting")
	}

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		listenAddr := fmt.Sprintf("%s:%d",
			s.conf.Config.ControlAPIConf.ListenAddr,
			s.conf.Config.ControlAPIConf.ListenPort,
		)

		if s.conf.Config.ControlAPIConf.UseSSL {
			errChan <- s.httpServer.StartTLS(
				listenAddr,
				s.conf.Config.ControlAPIConf.CertFilePath,
				s.conf.Config.ControlAPIConf.KeyFilePath,
			)
		} else {
			errChan <- s.httpServer.Start(listenAddr)
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-s.ctx.Done():
		err := s.httpServer.Shutdown(context.Background())
		<-errChan
		return err
	}
}
