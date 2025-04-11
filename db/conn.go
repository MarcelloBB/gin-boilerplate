package db

import (
	"fmt"

	"github.com/MarcelloBB/gin-boilerplate/config"
	"github.com/MarcelloBB/gin-boilerplate/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := config.LoadConfigIni("db", "host", "localhost").(string)
	port := config.LoadConfigIni("db", "port", 5432).(int)
	dbname := config.LoadConfigIni("db", "name", "postgres").(string)
	user := config.LoadConfigIni("db", "user", "postgres").(string)
	pwd := config.LoadConfigIni("db", "password", "1234").(string)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pwd, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)
	fmt.Println("Running automigration...")

	err = db.AutoMigrate(&model.Client{}, &model.User{}, &model.Product{})
	if err != nil {
		fmt.Println("Error running automigration:", err)
		panic(err)
	}

	return db, nil
}
