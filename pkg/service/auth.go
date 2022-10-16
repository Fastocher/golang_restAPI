package service

import (
	"crypto/sha1"
	"errors"
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

// Имплементирую логику описанную в сервисе
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	//Функция ParseWithClaims возвращает объект токена, в котором есть поле Claims типа interface
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		//промеряем метод подписи токена на HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signinKey), nil
	})

	if err != nil {
		return 0, err
	}
	// приведение Claims к структуре tokenClaims с проверкой
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type tokenClaims")
	}
	// при успешном парсинге возвращаем id пользователя
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
