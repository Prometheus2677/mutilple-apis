package repository

import (
	"log"

	"github.com/mr687/simple-go-rest-api/entity"
	"gorm.io/gorm"
)

// Define contract what's this repository can do
type UserRepository interface {
	Insert() (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint64) (*entity.User, error)
	FindByUsernameOrEmail(usernameOrEmail string) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	IsDuplicateEmail() bool
	IsDuplicateUsername() bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repo *userRepository) Insert(user *entity.User) (*entity.User, error) {
	// Doing something before insert
	err := user.BeforeInsert()
	if err != nil {
		log.Fatal(err)
	}

	err = repo.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) Save(user *entity.User) error {
	err := user.BeforeInsert()
	if err != nil {
		log.Fatal(err)
	}

	err = repo.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) FindByID(id uint64) (*entity.User, error) {
	var user *entity.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) FindByUsername(username string) (*entity.User, error) {
	var user *entity.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) FindByUsernameOrEmail(usernameOrEmail string) (*entity.User, error) {
	var user *entity.User
	err := repo.db.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userRepository) IsDuplicateEmail(email string) bool {
	_, err := repo.FindByEmail(email)
	return err == nil
}

func (repo *userRepository) IsDuplicateUsername(username string) bool {
	_, err := repo.FindByUsername(username)
	return err == nil
}

func (repo *userRepository) IsDuplicateEmailExcludeUser(email string, userId uint64) bool {
	var user *entity.User
	err := repo.db.Where("email = ? AND id != ?", email, userId).First(&user).Error
	return err == nil
}

func (repo *userRepository) IsDuplicateUsernameExcludeUser(username string, userId uint64) bool {
	var user *entity.User
	err := repo.db.Where("username = ? AND id != ?", username, userId).First(&user).Error
	return err == nil
}
