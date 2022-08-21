package handler

import (
	"context"
	"net/http"
	"time"

	srvconfig "github.com/k0tletka/sdfsd/internal/config"
	"github.com/k0tletka/sdfsd/internal/server/config"

	"github.com/labstack/echo/v4"
)

// ControlAPIHandler contains all handlers for Control API methods
type ControlAPIHandler struct {
	ctx  context.Context
	conf *config.ServerConfig
}

func NewControlAPIHandler(ctx context.Context, conf *config.ServerConfig) *ControlAPIHandler {
	return &ControlAPIHandler{
		ctx:  ctx,
		conf: conf,
	}
}

func (c *ControlAPIHandler) ServiceInfoHandler(ec echo.Context) error {
	return ec.JSON(http.StatusOK, ServiceInfoResponse{
		VersionNumber: srvconfig.VersionNumber,
		CommitHash:    srvconfig.CommitHash,
		BuildTime:     time.Unix(srvconfig.BuildTimeUnix, 0),
	})
}
