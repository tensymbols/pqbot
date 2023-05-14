package ports

import (
	"github.com/gin-gonic/gin"
	"vkbot/internal/app"
)

type Server struct {
	engine    *gin.Engine
	port      string
	accessKey string
}

func NewServer(port string, app app.App) *Server {
	gin.SetMode(gin.ReleaseMode)
	s := Server{
		engine: gin.Default(),
		port:   port,
	}
	Router(s.engine, app)
	return &s
}

func (s *Server) Run() {
	s.engine.Run()
}
