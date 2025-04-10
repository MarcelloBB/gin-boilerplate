package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "192.168.3.4"
	port   = 5432
	user   = "postgres"
	pwd    = "1234"
	dbname = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pwd, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
