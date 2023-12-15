# RESTful API - Golang

A simple RESTful API using [gin-gonic](https://github.com/gin-gonic/gin) for the web framework and [gorm](https://gorm.io) for the database ORM. This project is using Repository Pattern to separate the database and the business logic.

The project includes the following features:
* RESTful API endpoint
* Standard CRUD operations
* Data validation
* GEO IP lookup
* [JWT](https://jwt.io/)-based authentication
* Database migration
* Postman documentation and API testing

Note:
* The geoip info service is provided by [freegeoip.live](https://freegeoip.live) that may not be available when you try this project:(

## Getting Started
Please follow [the instructions](https://golang.org/doc/install) to install Go on your local machine. This project requires **Go 1.17 or higher+**.

We also need to install [Docker](https://www.docker.com/get-started) if you want to try the project without setting your own environment.

After all done, run the following commands to start the project:

```bash
# Clone this repo
git clone https://github.com/mr687/simple-go-rest-api.git

# Go to the project directory
cd simple-go-rest-api

# Install dependencies
go mod init

# Duplicate the .env.example file to .env and fill in your own values
cp .env.example .env

# Run postgreSQL database via Docker compose
docker compose up -d
# docker-compose up -d

# Start the rest server
go run main.go

# Server running default on http://localhost:8080
```

## API Documentation
The API documentation is hosted on [Postman Doc](https://documenter.getpostman.com/view/1838168/Uze1uio4).
