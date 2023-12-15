package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/response"
	"github.com/mr687/simple-go-rest-api/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Status:     false,
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       []string{err.Error()},
			})
		}

		c.Next()
	}
}
