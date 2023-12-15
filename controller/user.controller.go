package controller

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/dto"
	"github.com/mr687/simple-go-rest-api/helper"
	"github.com/mr687/simple-go-rest-api/repository"
	"github.com/mr687/simple-go-rest-api/response"
	"github.com/mr687/simple-go-rest-api/service"
)

func (s *Server) ChangeUsernameEmail(c *gin.Context) {
	var request dto.UpdateUsernameEmailRequest

	// Binding and validation
	if err := c.ShouldBind(&request); err != nil {
		response.ValidationError(c, err)
		return
	}

	// remove whitespace from username
	request.Username = strings.ReplaceAll(request.Username, " ", "")

	userId, err := service.GetTokenId(c)
	if err != nil || userId == 0 {
		response.Unauthorized(c)
		return
	}

	repo := repository.NewUserRepository(s.DB)
	user, err := repo.FindByID(userId)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	if repo.IsDuplicateEmailExcludeUser(request.Email, userId) {
		response.Conflict(c, "Email already exists")
		return
	}

	if repo.IsDuplicateUsernameExcludeUser(request.Username, userId) {
		response.Conflict(c, "Username already exists")
		return
	}

	user.Username = request.Username
	user.Email = request.Email
	err = repo.Save(user)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Ok(c, "Username and email changed", nil)
}

func (s *Server) ResetPassword(c *gin.Context) {
	userId, err := service.GetTokenId(c)
	if err != nil || userId == 0 {
		response.Unauthorized(c)
		return
	}

	repo := repository.NewUserRepository(s.DB)
	user, err := repo.FindByID(userId)
	if err != nil {
		response.Unauthorized(c)
		return
	}

	var request dto.ResetPasswordRequest

	// Binding and validation
	if err := c.ShouldBind(&request); err != nil {
		response.ValidationError(c, err)
		return
	}

	if !helper.ValidateHash(request.OldPassword, user.Password) {
		response.ValidationError(c, errors.New("wrong old password"))
		return
	}

	user.Password = request.NewPassword
	err = repo.Save(user)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Ok(c, "Password changed", nil)
}
