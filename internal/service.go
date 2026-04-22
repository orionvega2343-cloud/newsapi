package internal

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(email string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := User{Email: email, PasswordHash: string(hash)}
	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GenerateToken(userID int) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),                               // сюда id юзера, но как string
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // время истечения, jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *Service) Login(email string, password string) (string, error) {
	res, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := s.GenerateToken(res.Id)
	if err != nil {
		return "", err
	}
	return token, nil

}
