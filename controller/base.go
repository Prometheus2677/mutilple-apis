package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mr687/simple-go-rest-api/config"
	"github.com/mr687/simple-go-rest-api/entity"
	"github.com/mr687/simple-go-rest-api/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) Initialize(c config.DatabaseConfig) {
	var err error

	dbUri := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBName,
		c.DBPass,
	)

	s.DB, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to database")

		log.Fatal("error: ", err)
	} else {
		fmt.Println("Connected to database")
	}

	s.DB.AutoMigrate(
		&entity.User{},
		&entity.UserBalance{},
		&entity.UserBalanceHistory{},
		&entity.BankBalance{},
		&entity.BankBalanceHistory{},
	) // Database migration

	s.Router = gin.New()
	s.Router.Use(gin.Recovery(), middleware.LoggerMiddleware())

	s.InitializeRoutes()
}

func (s *Server) StartServer(c config.ServerConfig) {
	fmt.Printf("Listening to port %s\n", c.Port)

	addr := fmt.Sprintf(":%s", c.Port)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
