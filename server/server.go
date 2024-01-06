package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

type Server struct {
	Host    string
	Port    string
	Env     string
	Envvars []string
}

func (s *Server) getUrl() string {
	var protocol string
	if s.Env == "local" {
		protocol = "http"
	} else {
		protocol = "https"
	}

	return fmt.Sprintf("%s://%s:%s", protocol, s.Host, s.Port)
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

	// setup goth social providers
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), s.getUrl()+"/auth/google/callback"),
	)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", s.Host, s.Port)))
}

func NewServer() *Server {
	port := flag.String("p", "3000", "port")
	env := flag.String("e", "local", "env")
	flag.Parse()

	return &Server{
		Host: "0.0.0.0",
		Port: *port,
		Env:  *env,
		Envvars: []string{
			"GOOGLE_KEY",
			"GOOGLE_SECRET",
		},
	}
}
