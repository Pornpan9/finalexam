package customer

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/Pornpan9/finalexam/database"
	"strconv"
)

func (customer Customer) GetAll(conn *sql.DB) ([]Customer, error){
	tt := []Customer{}
	query := "SELECT id, name, email, status FROM customers"

	rows, err := conn.Query(query)
	if err != nil{
		return nil, err
	}

	for rows.Next(){
		var t Customer
		if err := rows.Scan(&t.ID, &t.Name, &t.Email, &t.Status); err != nil{
			return nil, err			
		}
		tt = append(tt, t)
	}
	return tt, err
}

func (customer Customer) GetByID(conn *sql.DB) (Customer, error){

	query := `
		SELECT 	id, name, email, status 
		FROM 	customers 
		where 	id = $1;
	`
	row := conn.QueryRow(query, customer.ID)

	var t Customer
	if err := row.Scan(&t.ID, &t.Name, &t.Email, &t.Status); err != nil{
		return t, err			
	}

	return t, nil
}

func GetHandler(c *gin.Context)  {

	conn, err := database.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	defer conn.Close()

	t := Customer{}
	tt, err := t.GetAll(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, tt)
}

func GetByIDHandler(c *gin.Context)  {

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



	t, err = t.GetByID(conn)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)

}