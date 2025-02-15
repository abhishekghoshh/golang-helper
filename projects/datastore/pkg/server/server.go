package server

import (
	"net/http"

	"github.com/labstack/echo"
)

type Server interface {
	Init(e *echo.Echo)
}

type OthersApi struct {
}

func OthersApis() Server {
	return &OthersApi{}
}

func (api *OthersApi) Init(e *echo.Echo) {
	e.Static("/", "static")
	e.GET("/health", api.healthCheck)

}

// Health check API
func (api *OthersApi) healthCheck(c echo.Context) error {
	data := map[string]string{
		"status": "up and running",
	}
	return c.JSON(http.StatusOK, data)
}
