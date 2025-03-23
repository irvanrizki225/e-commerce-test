package services

import (
	"e-commerce/objects"
	"e-commerce/models"
)


func FetchProduct(params objects.Params) ([]objects.Product, error) {
	var products []objects.Product

	query := db.Model(&models.Product{})

	if params.Search != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+params.Search+"%")
	}

	err := query.Find(&products).Error
	return products, err
}

func GetProductByID(id int) (objects.Product, error) {
	var product objects.Product
	err := db.Where("id = ?", id).First(&product).Error
	return product, err
}

func CreateCartProduct(product_id uint, user_id uint, quantity int) (bool, error) {
	var product objects.Product

	err := db.Where("id = ?", product_id).First(&product).Error
	if err != nil {
		return false, err
	}

	if product.Stock < quantity {
		return false, nil
	}

	product.Stock -= quantity
	err = db.Model(&models.Product{}).Where("id = ?", product_id).Update("stock", product.Stock).Error
	if err != nil {
		return false, err
	}

	card := models.Card{
		ProductID: product_id,
		UserID:    user_id,
		Quantity:  quantity,
	}

	err = db.Create(&card).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckOrder(user_id uint, cart_id uint, payment_method string) (bool, error) {
	var order models.Order

	err := db.Where("user_id = ?", user_id).First(&order).Error
	if err != nil {
		return false, err
	}

	order.PaymentMethod = payment_method

	err = db.Save(&order).Error
	if err != nil {
		return false, err
	}

	err = db.Where("id = ?", cart_id).Delete(&models.Card{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetOrder(order_id uint) ([]objects.Order, error) {
	var orders []objects.Order

	err := db.Where("id = ?", order_id).Find(&orders).Error
	return orders, err
}