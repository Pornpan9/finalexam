package main

import(
	"github.com/Pornpan9/finalexam/customer"
	"github.com/Pornpan9/finalexam/database"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main()  {
	
	database.InitDB()
	r := setupRouter()
	r.Run(":2019")
}

func authMiddleware(c *gin.Context){
	fmt.Println("Hello")
	token := c.GetHeader("Authorization")
	fmt.Println("token:", token)

	if token != "token2019"{
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}

	c.Next()
	fmt.Println("Goodbye!!!!")
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(authMiddleware)

	api := r.Group("/")
	api.GET("/customers", customer.GetHandler)
	api.POST("/customers", customer.CreateHandler)
	api.PUT("/customers/:id", customer.UpdateHandler)
	api.DELETE("/customers/:id", customer.DeleteHandler)
	api.GET("/customers/:id", customer.GetByIDHandler)	

	return r
}
