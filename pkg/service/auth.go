package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Fastocher/restapp"
	"github.com/Fastocher/restapp/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const salt = "maifoguefsnaps235ti4p"
const tokenTTL = 12 * time.Hour
const signinKey = "adfhadofhiwrn13t"

type AuthService struct {
	repo repository.Authorization
}

// Дополнение стандартного claims
// для добавления поля с id нашего пользователя
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user restapp.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//получаем пользователя из базы данных
	user, err := s.repo.GetUser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}
	//заменили стандартный Claims на наш + стандартный Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signinKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
