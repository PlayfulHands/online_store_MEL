package main

import (
	"Lesson/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Маршруты для аутентификации и регистрации
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	// Маршруты для каталога товаров и корзины
	r.GET("/catalog", handlers.GetCatalog)
	r.GET("/cart", handlers.GetCart)

	r.Run(":8080") // Запуск веб-сервера на порту 8080
	// Создаем экземпляр маршрутизатора Gin
	router := gin.Default()

	// Настройка маршрута для эндпоинта каталога
	router.GET("/catalog", handlers.GetCatalog)

	// Запуск сервера на порту 8080
	router.Run(":8080")
}
