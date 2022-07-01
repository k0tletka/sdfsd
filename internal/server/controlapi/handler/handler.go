package handler

import (
	"context"
	"io"
	"net/http"

	"github.com/k0tletka/SDFS/internal/server/config"
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

func (c *ControlAPIHandler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
