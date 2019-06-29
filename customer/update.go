package customer

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/Pornpan9/finalexam/database"
	"strconv"
)

func (customer Customer) Update(conn *sql.DB) error{

	query := `
		UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;
	`;

	stmt, err := conn.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email, customer.Status)
	
	return err
}

func UpdateHandler(c *gin.Context)  {
	
	t := Customer{}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))//get param on url
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	t.ID = id

	conn, err := database.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	defer conn.Close()

	err = t.Update(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}