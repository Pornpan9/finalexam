package customer

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/Pornpan9/finalexam/database"
)

func (customer Customer) Insert(conn *sql.DB) (Customer,error){

	query := `
		INSERT INTO customers (name, email, status)
		VALUES ($1, $2, $3) RETURNING id
	`;

	row := conn.QueryRow(query, customer.Name, customer.Email, customer.Status)
	err := row.Scan(&customer.ID)

	return customer, err
}

func CreateHandler(c *gin.Context)  {
	t := Customer{}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	conn, err := database.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	defer conn.Close()

	t, err = t.Insert(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, t)
}