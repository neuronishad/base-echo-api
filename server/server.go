package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Host string
	Port string
}

func (s *Server) Start() {
	e := echo.New()
	e.HideBanner = true

	s.addRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", s.Host, s.Port)))
}

func NewServer() *Server {
	return &Server{
		Host: "0.0.0.0",
		Port: "3000",
	}
}
