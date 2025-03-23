package routes

import (
	"e-commerce/controllers"
	"e-commerce/auth"
	"e-commerce/helpers"
	
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

)

var (
	authServices auth.Service = auth.NewService()
)

func SetRouter() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
	}
	r.Use(cors.New(config))

	r.MaxMultipartMemory = 8 << 20

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/products", authMiddleware(authServices), controllers.GetAllProduct)
	r.GET("/products/:id", authMiddleware(authServices), controllers.GetProductByID)

	r.POST("/carts", authMiddleware(authServices), controllers.AddCartProduct)

	r.POST("/orders", authMiddleware(authServices), controllers.CheckOrder)

	err := r.Run(":9060")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":9060", r))
}

func authMiddleware(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helpers.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helpers.APIResponse("Invalid token claims", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user_id := claim["user_id"].(string)

		//validation user
		validation := helpers.ValidateUser(user_id)
		if validation != "success" {
			response := helpers.APIResponse(validation, http.StatusBadRequest, "error", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		c.Set("currentUser", user_id)
	}
}
