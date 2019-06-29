package main

import(
	"github.com/Pornpan9/finalexam/customer"
	"github.com/Pornpan9/finalexam/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	
	database.InitDB()
	r := setupRouter()
	r.Run(":2019")
}

func authMiddleware(c *gin.Context){
	token := c.GetHeader("Authorization")

	if token != "token2019"{
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}

	c.Next()
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(authMiddleware)

	r.GET("/customers", customer.GetHandler)
	r.POST("/customers", customer.CreateHandler)
	r.PUT("/customers/:id", customer.UpdateHandler)
	r.DELETE("/customers/:id", customer.DeleteHandler)
	r.GET("/customers/:id", customer.GetByIDHandler)	

	return r
}
