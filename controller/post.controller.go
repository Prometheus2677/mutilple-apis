package controller

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/mr687/simple-go-rest-api/dto"
	"github.com/mr687/simple-go-rest-api/entity"
	"github.com/mr687/simple-go-rest-api/repository"
	"github.com/mr687/simple-go-rest-api/response"
)

func (s *Server) CreatePost(c *gin.Context) {
	var createRequest dto.CreateRequest

	if err := c.ShouldBind(&createRequest); err != nil {
		response.ValidationError(c, err)
		return
	}

	createRequest.Title = strings.ReplaceAll(createRequest.Title, " ", "")

	newPost := &entity.Post{}
	_ = smapping.FillStruct(&newPost, smapping.MapFields(&createRequest))

	repo := repository.NewPostRepository(s.DB)

	postCreated, err := repo.CreatePost(newPost)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}
	response.Created(c, "Post Created", postCreated)
}

func (s *Server) GetPost(c *gin.Context) {
	idString := c.Param("id")
	repo := repository.NewPostRepository(s.DB)
	id, err := strconv.ParseUint(idString, 10, 64)
	postFound, err := repo.FindById(id)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}
	response.Created(c, "Post Found", postFound)
}
