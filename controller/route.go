package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/middleware"
	"github.com/mr687/simple-go-rest-api/response"
)

func (s *Server) InitializeRoutes() {
	// Index Route
	s.Router.GET("/", s.Index)

	// API Routes
	api := s.Router.Group("/api/v1")
	{
		// Auth Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.Register)
			auth.POST("/login", s.Login)
			auth.POST("/logout", middleware.AuthMiddleware(), s.Logout)
		}

		// User Routes
		user := api.Group("/user", middleware.AuthMiddleware())
		{
			auth.PUT("/reset-password", middleware.AuthMiddleware(), s.ResetPassword)
			user.PUT("/change-username-email", s.ChangeUsernameEmail)
		}

		// User Balance Routes
		balance := api.Group("/balance", middleware.GeoIpMiddleware())
		{
			balance.GET("/current", s.GetBalance)
			balance.POST("/topup", s.AddBalance)
			balance.POST("/send", s.SendBalance)
		}

		// Post
		post := api.Group("/post", middleware.AuthMiddleware())
		{
			post.POST("/create", s.CreatePost)
			post.GET("/get/:id", s.GetPost)
		}
	}

	// Otherwise, just show Not Found 404
	s.Router.NoRoute(func(c *gin.Context) {
		response.NotFound(c)
	})

}
