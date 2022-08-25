package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/vicfntm/splitshit2/src/models"
	"github.com/vicfntm/splitshit2/src/repositories"
)

type AuthService struct {
	repo     repositories.Authorization
	JWT_SIGN string
	JWT_SALT string
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo *repositories.Repositories) *AuthService {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env file loading failed: %s", err.Error())
	}

	return &AuthService{
		repo:     repo.Authorization,
		JWT_SIGN: os.Getenv("JWT_SIGN"),
		JWT_SALT: os.Getenv("JWT_SALT"),
	}
}

func (as *AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(as.JWT_SIGN), nil

	})
	if err != nil {

		return 0, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, err
	}

}

func (as *AuthService) CreateCustomer(user models.Customer) (int, error) {
	user.Password = generatePasswordHash(user.Password, as)
	return as.repo.CreateUser(user)
}

func (as *AuthService) LoginUser(user models.Customer) (string, error) {
	dbUser, error := as.repo.GetUserFromDb(user.Login, generatePasswordHash(user.Password, as))
	if error != nil {
		log.Printf("no user found")

		return "", error
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, dbUser.Id,
	})

	tokenByte, error := token.SignedString([]byte(as.JWT_SIGN))

	if error != nil {
		log.Printf("token creation failed")
	}

	if error != nil {
		log.Printf("token sending failed")
	}

	return tokenByte, error
}

func (as *AuthService) FindCustomer(id interface{}) (models.Customer, error) {
	return as.repo.GetCustomerById(id)
}

func generatePasswordHash(password string, as *AuthService) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(as.JWT_SALT)))
}

func (as *AuthService) AssignRole(id int, role string) (models.RoleCustomer, error) {

	return as.repo.AssignRole(id, role)
}
