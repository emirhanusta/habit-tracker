package server

import (
	"github.com/gofiber/fiber/v2"
	"habit-tracker/configuration"
)

type Server struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *Server {
	return &Server{app: app}
}

func (s *Server) StartServer() {
	err := s.app.Listen(":" + configuration.Port)
	if err != nil {
		return
	}
}
