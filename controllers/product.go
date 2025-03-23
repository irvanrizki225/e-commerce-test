package controllers

import (
	"e-commerce/helpers"
	"e-commerce/objects"
	"e-commerce/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	var params objects.Params

	if err := c.ShouldBindQuery(&params); err != nil {
		response := helpers.APIResponse("failed get all product", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := services.FetchProduct(params)
	if err != nil {
		response := helpers.APIResponse("failed get all product", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("success get all product", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func GetProductByID(c *gin.Context) {
	var id int	
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		response := helpers.APIResponse("failed get product by id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := services.GetProductByID(id)
	if err != nil {
		response := helpers.APIResponse("failed get product by id", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("success get product by id", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func AddCartProduct(c *gin.Context) {
	var cartProduct objects.Cart

	if err := c.ShouldBindJSON(&cartProduct); err != nil {
		response := helpers.APIResponse("failed add product to cart", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cart, err := services.CreateCartProduct(cartProduct.ProductID, cartProduct.UserID, cartProduct.Quantity)
	if err != nil {
		response := helpers.APIResponse("failed add product to cart", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if !cart {
		response := helpers.APIResponse("failed add product to cart", http.StatusBadRequest, "error", "Product out of stock")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("success add product to cart", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func CheckOrder(c *gin.Context) {
	var orderRequest objects.CheckOut

	user_ids, _ := c.Get("currentUser")
	user_id := uint(user_ids.(float64))

	card_ids, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		response := helpers.APIResponse("failed get product by id", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	CartID := uint(card_ids)

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		response := helpers.APIResponse("failed check out", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := services.CheckOrder(user_id, CartID, orderRequest.PaymentMethod)
	if err != nil {
		response := helpers.APIResponse("failed check out", http.StatusInternalServerError, "error", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if !order {
		response := helpers.APIResponse("failed check out", http.StatusBadRequest, "error", "Product out of stock")
		c.JSON(http.StatusBadRequest, response)
		return
	}
}