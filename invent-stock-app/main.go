package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	//load .env file
	err := godotenv.Load("../.env")
	if err != nil {
	  log.Fatal("Error loading .env file ", err)
	}

	//db connection
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Error connecting to database ", err)
	}
	defer db.Close()

	fmt.Println("Connection to db successful ", db)
}
