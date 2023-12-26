package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Host    string
	Port    string
	Envvars []string
}

func (s *Server) Start() {
	e := echo.New()
	e.HideBanner = true

	s.addRoutes(e)

	// ensure Envvars
	for _, envV := range s.Envvars {
		if val := os.Getenv(envV); val == "" {
			fmt.Printf("Environment Variable '%s' probably not set\n", envV)
		}
	}

	e.Static("/static", "assets")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", s.Host, s.Port)))
}

func NewServer() *Server {
	return &Server{
		Host: "0.0.0.0",
		Port: "3000",
	}
}
