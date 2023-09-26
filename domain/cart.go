package domain

type CartItem struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

func AddToCart(cart *Cart, product Product, quantity int) {
	// Создаем элемент корзины на основе продукта и количества
	cartItem := CartItem{
		ProductID: product.ID, // Важно использовать поле ID, а не Id
		Quantity:  quantity,
		Price:     product.Price,
	}

	// Добавляем элемент корзины в список
	cart.Items = append(cart.Items, cartItem)
}
