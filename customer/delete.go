package customer

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/Pornpan9/finalexam/database"
	"strconv"
)


func (customer Customer) Delete(conn *sql.DB) error{

	query := `
		DELETE FROM customers WHERE id=$1;
	`;

	stmt, err := conn.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(customer.ID)
	
	return err
}


func DeleteHandler(c *gin.Context)  {

	t := Customer{}
	
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

	err = t.Delete(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	// c.JSON(http.StatusOK, t)
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
