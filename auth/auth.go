package auth

import (
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"


	"e-commerce/objects"
)

type Service interface {
	GenerateToken(user objects.User) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
	ValidatePassword(userPassword string, RequestPassword string) error
}

type jwtService struct{}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewService() Service {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(user objects.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	jwtClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirationTime,
	}

	//generate token with expires time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString(SECRET_KEY)
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return SECRET_KEY, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
}

func (s *jwtService) ValidatePassword(userPassword string, RequestPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(RequestPassword))
}