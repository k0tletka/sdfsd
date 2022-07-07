package controlapi

import (
	"github.com/k0tletka/SDFS/internal/server/controlapi/handler"
	"github.com/labstack/echo/v4"
)

func (c *ControlAPI) registerRoutes(handlers *handler.ControlAPIHandler, e *echo.Echo) {
	apiv1 := e.Group("/api/v1")
	apiv1.GET("/serverInfo", handlers.ServiceInfoHandler)
}
