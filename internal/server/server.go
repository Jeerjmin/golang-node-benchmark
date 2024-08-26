package server

import (
	"fmt"
	"go-auth/api"
	"go-auth/internal/config"
	database "go-auth/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config *config.Config
	Router *gin.Engine
}

func NewServer(config *config.Config, db *database.Database) *Server {
	router := api.RegisterRoutes(db)
	return &Server{
		Config: config,
		Router: router,
	}
}

func (s *Server) Start() {
	fmt.Printf("Starting server at port %s...\n", s.Config.Port)

	if err := s.Router.Run(":" + s.Config.Port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
