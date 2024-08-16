package server

import "github.com/gofiber/fiber/v2"

type Server struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *Server {
	return &Server{app: app}
}

func (s *Server) StartServer() {
	err := s.app.Listen(":8080")
	if err != nil {
		return
	}
}
