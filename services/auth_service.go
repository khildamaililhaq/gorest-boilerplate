package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/khildamaililhaq/gorest-boilerplate/dto"
	"github.com/khildamaililhaq/gorest-boilerplate/models"
	"os"
	"time"
)

type AuthService interface {
	Login(user dto.LoginDTO) models.User
	Register(user dto.RegisterDTO) models.User
}

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService  {
	return &jwtService{
		issuer: "khildamaililhaq",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string{
	secretKey := os.Getenv("SECRET_KEY")

	if secretKey == "" {
		secretKey = "uta"
	}

	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1,0,0).Unix(),
			Issuer: "Khilda Maililhaq",
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error){
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}