package db

import (
	"database/sql"
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/config"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	host := config.LoadConfigIni("db", "host", "localhost").(string)
	port := config.LoadConfigIni("db", "port", 5432).(int)
	dbname := config.LoadConfigIni("db", "name", "postgres").(string)
	user := config.LoadConfigIni("db", "user", "postgres").(string)
	pwd := config.LoadConfigIni("db", "password", "1234").(string)

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
