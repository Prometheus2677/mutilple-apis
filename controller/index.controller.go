package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/response"
)

func (s *Server) Index(c *gin.Context) {
	response.Ok(c, "Simple API - Privy BE Test", nil)
}
