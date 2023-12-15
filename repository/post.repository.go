package repository

import (
	"github.com/mr687/simple-go-rest-api/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost() (*entity.Post, error)
	FindByTitle(title string) (*entity.Post, error)
	UpdatePost(id uint64) (*entity.Post, error)
	DeletePost(id uint64) bool
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db}
}

func (repo *postRepository) CreatePost(post *entity.Post) (*entity.Post, error) {
	err := repo.db.Create(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (repo *postRepository) FindByTitle(title string) (*entity.Post, error) {
	var post *entity.Post
	err := repo.db.Where("title = ?", title).First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}
