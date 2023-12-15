package controller

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/mr687/simple-go-rest-api/dto"
	"github.com/mr687/simple-go-rest-api/entity"
	"github.com/mr687/simple-go-rest-api/helper"
	"github.com/mr687/simple-go-rest-api/repository"
	"github.com/mr687/simple-go-rest-api/response"
	"github.com/mr687/simple-go-rest-api/service"
)

func (s *Server) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest

	// Binding and validation
	if err := c.ShouldBind(&registerRequest); err != nil {
		response.ValidationError(c, err)
		return
	}

	// remove whitespace from username
	registerRequest.Username = strings.ReplaceAll(registerRequest.Username, " ", "")

	newUser := &entity.User{}
	_ = smapping.FillStruct(&newUser, smapping.MapFields(&registerRequest))

	repo := repository.NewUserRepository(s.DB)

	if repo.IsDuplicateEmail(newUser.Email) {
		response.Conflict(c, "Email already exists")
		return
	}

	if repo.IsDuplicateUsername(newUser.Username) {
		response.Conflict(c, "Username already exists")
		return
	}

	userCreated, err := repo.Insert(newUser)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Created(c, "User created", userCreated)
}

type LoginResponse struct {
	User      *entity.User `json:"user"`
	Token     string       `json:"token"`
	ExpiredAt int64        `json:"expired_at"`
}

func (s *Server) Login(c *gin.Context) {
	var credential dto.LoginCredential

	// Binding and validation
	if err := c.ShouldBind(&credential); err != nil {
		response.ValidationError(c, err)
		return
	}

	repo := repository.NewUserRepository(s.DB)

	// check if user exists
	existsUser, err := repo.FindByUsernameOrEmail(credential.Account)
	if err != nil {
		response.ValidationError(c, errors.New("user not found"))
		return
	}
	// check if the user password is correct
	if !helper.ValidateHash(credential.Password, existsUser.Password) {
		response.ValidationError(c, errors.New("wrong password"))
		return
	}

	// Generate auth token
	token, expiredAt := service.GenerateToken(existsUser.Id)

	response.Ok(c, "Login success", &LoginResponse{
		User:      existsUser,
		Token:     token,
		ExpiredAt: expiredAt,
	})
}

func (s *Server) Logout(c *gin.Context) {
	// Return a success response
	response.Ok(c, "Logged out successfully", nil)
}
