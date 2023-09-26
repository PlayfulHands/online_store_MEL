package domain

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)
import "errors"

var ErrUnauthorized = errors.New("Unauthorized")
var users []User
var SecretKey = []byte("your_secret_key_here")

type User struct {
	Id       int32
	Username string
	Password string
}

func Authenticate(username, password string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {
				return &user, nil
			}
		}
	}
	return nil, ErrUnauthorized
}

func UserExists(username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func CreateUser(user User) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	users = append(users, user)
}

func GenerateToken(user *User) (string, error) {
	// Создание токена
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Пример: токен действителен 24 часа

	// Подпись токена с использованием секретного ключа
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
