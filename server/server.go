package server

import "github.com/gin-gonic/gin"

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	router := gin.Default()

	apiGroups := router.Group("/api/v1")

	authGroup := apiGroups.Group("/auth")
	{
		authGroup.POST("/token")
	}

	videoGroup := apiGroups.Group("video")
	{
		videoGroup.Use(authenticate())
		
		videoGroup.GET("/share")
		videoGroup.POST("/upload")
		videoGroup.PATCH("/edit")
	}

	return nil
}
