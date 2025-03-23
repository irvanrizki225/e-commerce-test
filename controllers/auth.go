package controllers

import (
	"e-commerce/auth"
	"e-commerce/helpers"
	"e-commerce/objects"
	"e-commerce/services"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login objects.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input jika struct memiliki tag validasi
	if err := validator.New().Struct(login); err != nil {
		validate := helpers.FormatValidatorError(err)
		response := helpers.APIResponse("failed login", http.StatusBadRequest, "error", validate)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := services.GetUserByEmail(login.Email)
	if err != nil {
		response := helpers.APIResponse("failed login", http.StatusUnauthorized, "error", "Invalid credentials")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	authService := auth.NewService()
	// Cek password
	if err := authService.ValidatePassword(user.Password, login.Password); err != nil {
		response := helpers.APIResponse("failed login", http.StatusUnauthorized, "error", "Invalid credentials")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// Generate token
	token, err := authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Update token user (jika perlu)
	if err := services.UpdateUserToken(user.ID, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update token"})
		return
	}

	// Berikan respons sukses dengan token
	response := helpers.APIResponse("success login", http.StatusOK, "success", gin.H{"token": token})
	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context) {
	var register objects.Register

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// Validasi input jika struct memiliki tag validasi
	if err := validator.New().Struct(register); err != nil {
		validate := helpers.FormatValidatorError(err)
		response := helpers.APIResponse("failed register", http.StatusBadRequest, "error", validate)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := services.CreateUser(register)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	response := helpers.APIResponse("success register", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}