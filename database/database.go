package database

import(
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

var dbURL string

func InitDB(){
	dbURL = os.Getenv("DATABASE_URL")
	if len(dbURL) == 0 {
		log.Fatal("Environment vaiable DATABASE_URL is empty")
	}

	db, err := sql.Open("postgres",dbURL)
	if err != nil {
		log.Fatal("Can't connect db",err.Error())
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS customers(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			status TEXT
	);
	`;

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table fatal error",err.Error())
	}
}

func Connect()(*sql.DB, error){
	db, err := sql.Open("postgres", dbURL)
	return db, err
}