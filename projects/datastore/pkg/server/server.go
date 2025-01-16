package server

import "github.com/labstack/echo"

type Server interface {
	Init(e *echo.Echo)
}
