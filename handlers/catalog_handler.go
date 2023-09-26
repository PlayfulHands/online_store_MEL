package handlers

import (
	"Lesson/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCatalog(c *gin.Context) {
	// Здесь должна быть логика получения списка товаров.
	// В данном примере мы создадим фиктивный список товаров.
	products := []domain.Product{
		{
			ID:          1,
			Name:        "Product 1",
			Description: "Description for Product 1",
			Price:       19.99,
		},
		{
			ID:          2,
			Name:        "Product 2",
			Description: "Description for Product 2",
			Price:       29.99,
		},
		// Добавьте другие товары по аналогии
	}

	// Отправляем список товаров в формате JSON клиенту
	c.JSON(http.StatusOK, products)
}

// Другие функции обработки запросов, если есть
func GetCartCatalog(c *gin.Context) {
	// Код для получения товаров в корзине
}
