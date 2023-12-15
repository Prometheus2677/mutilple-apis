package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtCustomClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

func GetSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "verylongsecretkey"
	}
	return secretKey
}

func GenerateToken(userId uint64) (string, int64) {
	// Define expire time for 1 month
	today := time.Now()
	expiresAt := today.Add(time.Minute * 30).Unix() // Token expires after 30 mins

	claims := &jwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "somethinghere",
			IssuedAt:  today.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(GetSecretKey()))
	if err != nil {
		return "", 0
	}

	return signedToken, expiresAt
}

func VerifyToken(c *gin.Context) error {
	token, err := ParseToken(c)
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwtCustomClaims); ok {
		fmt.Println(claims)
	}

	return nil
}

func ParseToken(c *gin.Context) (*jwt.Token, error) {
	tokenString, err := GetToken(c)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetSecretKey()), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err := errors.New("invalid signature")
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		err := errors.New("token is invalid")
		return nil, err
	}

	return token, nil
}

func GetTokenId(c *gin.Context) (uint64, error) {
	token, _ := ParseToken(c)

	if token == nil {
		err := errors.New("token is empty")
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, nil
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		return 0, err
	}

	return uint64(userId), nil
}

func GetToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	println(authHeader)

	if len(authHeader) == 0 {
		err := errors.New("authorization header is empty")
		return "", err
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		err := errors.New("authorization header is invalid format")
		return "", err
	}

	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		err := errors.New("authorization header is unsupported")
		return "", err
	}

	token := fields[1]
	if len(token) == 0 {
		err := errors.New("token is empty")
		return "", err
	}
	return token, nil
}
