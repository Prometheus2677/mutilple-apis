package response

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type EmptyData struct{}

func Ok(c *gin.Context, message string, data interface{}) {
	MakeResponse(c, true, message, http.StatusOK, data)
}

func Created(c *gin.Context, message string, data interface{}) {
	MakeResponse(c, true, message, http.StatusCreated, data)
}

func Conflict(c *gin.Context, message string) {
	MakeResponse(c, false, message, http.StatusConflict, nil)
}

func BadRequest(c *gin.Context, message string) {
	MakeResponse(c, false, message, http.StatusBadRequest, nil)
}

func NotFound(c *gin.Context) {
	MakeResponse(c, false, "Not Found", http.StatusNotFound, nil)
}

func ValidationError(c *gin.Context, data interface{}) {
	MakeResponse(c, false, "Validation Error", http.StatusBadRequest, data)
}

func Unauthorized(c *gin.Context) {
	MakeResponse(c, false, "Unauthorized", http.StatusUnauthorized, nil)
}

func ServerError(c *gin.Context, data interface{}) {
	MakeResponse(c, false, "Server Error", http.StatusInternalServerError, data)
}

func MakeResponse(c *gin.Context, status bool, message string, statusCode int, data interface{}) {
	if !status && data != nil {
		splittedErrors := strings.Split(fmt.Sprint(data), "\n")
		data = splittedErrors
	}

	if data == nil || reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface()) {
		c.JSON(statusCode, Response{
			Status:     status,
			StatusCode: statusCode,
			Message:    message,
		})
	} else {
		c.JSON(statusCode, Response{
			Status:     status,
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
		})
	}
}
