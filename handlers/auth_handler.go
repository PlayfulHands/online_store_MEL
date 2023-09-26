package handlers

import (
	"Lesson/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	// Обработка аутентификации пользователя
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Проверка аутентификации (здесь нужно реализовать логику аутентификации)
	user, err := domain.Authenticate(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Создание и отправка токена (JWT) для аутентифицированного пользователя
	// (необходимо реализовать функциональность создания токена)
	token, err := domain.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	// Обработка регистрации нового пользователя
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Проверка наличия пользователя с таким именем (здесь нужно реализовать логику)
	if domain.UserExists(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Создание нового пользователя и сохранение его данных
	newUser := domain.User{
		Username: username,
		Password: password,
	}
	domain.CreateUser(newUser)

	// Генерация токена для нового пользователя
	token, err := domain.GenerateToken(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful", "token": token})
}
