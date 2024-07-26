package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func displayMenu () {

	menu := map[int]string{
		1: "Add Product",
		2: "Change Stock",
		3: "Add Staff", 
		4: "Generate Sales Report",
		5: "Exit",
	}
	fmt.Println("==========================================")
	fmt.Println("Welcome to Store Management CLI Interface")
	fmt.Println("==========================================")
	fmt.Println("")
	fmt.Println("Please select menu from below...")
	fmt.Println("")
	for i , v := range menu{
		fmt.Printf("%d. %s\n", i, v)
	}
}

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

	fmt.Println("Connection to db successful ")
	fmt.Println("")

	//init bufio reader
	reader := bufio.NewReader(os.Stdout)

	//run apps
	menuLoop: for{
		displayMenu()
		fmt.Printf("Answer (1/2/3/4/5): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input ", err)
		}
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			break menuLoop;
		case "2":
			break menuLoop;
		case "3":
			break menuLoop;
		case "4":
			break menuLoop;
		case "5":
			break menuLoop;
		}
	}



}
