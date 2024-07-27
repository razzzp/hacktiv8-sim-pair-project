package main

import (
	"bufio"
	"fmt"
	"invent-stock-app/repo"
	"log"
	"net/mail"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


type ProductModifStockParam struct {
	ProductName   string
	StockNumToAdd int
}

func displayMenu() {

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
	for i, v := range menu {
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
	db, err := repo.CreateDBInstance()
	if err != nil {
		log.Fatal("Error connecting to database ", err)
	}
	defer db.Close()

	fmt.Println("Connection to db successful ")
	fmt.Println("")

	// create repos
	productRepo := repo.CreateProductRepo(db)
	staffRepo := repo.CreateStaffRepo(db)

	//init bufio reader
	reader := bufio.NewReader(os.Stdout)

	//run apps
menuLoop:
	for {
		displayMenu()
		fmt.Printf("Answer (1/2/3/4/5): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input ", err)
		}
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			product := getAddProdParam(reader)
			if isZeroValue(product) {
				continue
			}
			productRepo.AddProduct(&product)
			fmt.Printf("Successfully add %s (%.2f) with qty of %d unit\n", product.Name, product.Price, product.Stock)
			fmt.Println("")
		case "2":
			RunProducStockModif(reader, productRepo)
			fmt.Println("")
		case "3":
			staff := getAddStaffParam(reader)
			if isZeroValue(staff) {
				continue
			}
			staffRepo.AddStaff(&staff)
			fmt.Printf("Successfully add %s as %s staff\n", staff.Name, staff.Position)
			fmt.Println("")
		case "4":
			break menuLoop
		case "5":
			break menuLoop
		default:
			fmt.Println("Please enter a valid option.")
		}
	}
}

//function to make user able to select option wether backt reinputting or to main menu
func selectOption [K repo.Staff | repo.Product ] (reader *bufio.Reader, function func(*bufio.Reader) K) K {
	fmt.Printf("Would you like to reinput data (y/n): ")
	resp, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading response input", err)
	}
	resp = strings.TrimSpace(resp)
	var zero K
	switch resp {
	case "y":
		return function(reader)
	case "n":
		return zero
	default:
		fmt.Println("Input invalid, please reinput response")
		selectOption(reader, function)
	}
	return zero

}
//function to make user back to main menu
func isZeroValue[K repo.Staff | repo.Product](value K) bool {
	var zero K
	return value == zero
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getAddStaffParam(reader *bufio.Reader) repo.Staff {
	//get product name
	fmt.Printf("Please insert staff name: ")
	staffName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading staff name input ", err)
	}
	staffName = strings.TrimSpace(staffName)
	if staffName == "" {
		log.Println("Staff name should not be empty")
		return selectOption(reader, getAddStaffParam)
	}

	//get staff email
	fmt.Printf("Please insert staff email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading product email input ", err)
	}
	email = strings.TrimSpace(email)
	if email == "" {
		log.Println("Staff email should not be empty")
		return selectOption(reader, getAddStaffParam)
	}
	//validate email
	if res := isValidEmail(email); !res {
		log.Println("Invalid Email Format")
		return selectOption(reader, getAddStaffParam)
	}

	//get staff position
	fmt.Printf("Please insert product position: ")
	position, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading staff position input ", err)
	}
	position = strings.TrimSpace(position)
	if position == "" {
		log.Println("Staff position should not be empty")
		return selectOption(reader, getAddStaffParam)
	}

	return repo.Staff{
		Name:     staffName,
		Email:    email,
		Position: position,
	}
}

func getAddProdParam(reader *bufio.Reader) repo.Product {
	//get product name
	fmt.Printf("Please insert product name: ")
	prodName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading product name input ", err)
	}
	prodName = strings.TrimSpace(prodName)
	if prodName == "" {
		log.Println("Product name should not be empty")
		return selectOption(reader, getAddProdParam)

	}

	//get product price
	fmt.Printf("Please insert product price: ")
	price, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading product price input ", err)
	}
	price = strings.TrimSpace(price)
	priceFlt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Price should be a number")
		return selectOption(reader, getAddProdParam)
	}

	//get product stock
	fmt.Printf("Please insert product stock: ")
	stock, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading product stock input ", err)
	}
	stock = strings.TrimSpace(stock)
	stockInt, err := strconv.Atoi(stock)
	if err != nil {
		log.Println("Stock should be a number")
		return selectOption(reader, getAddProdParam)
		
	}
	if stockInt < 0 {
		log.Println("Stock should be a positive number")
		return selectOption(reader, getAddProdParam)
	}
	return repo.Product{
		Name: prodName,
		Price: priceFlt,
		Stock: stockInt,
	}
}

// prompts user for product to modify stock
func RunProducStockModif(reader *bufio.Reader, productRepo repo.ProductRepo) {
	// prompt product name
	fmt.Println("Please enter product name to modify stock:")
	input, err := reader.ReadString('\n')
	if err != nil {
		// something went wrong exit program
		log.Fatal(err)
	}

	productName := strings.TrimSpace(input)
	if productName == "" {
		fmt.Println("Product name cannot be empty.")
		return
	}

	// check that product exists in DB
	existingProduct, err := productRepo.GetProductByName(productName)
	if err != nil || existingProduct == nil {
		fmt.Printf("Product '%s' does not exist.", productName)
		return
	}

	// prompt num of stock to add or reduce
	//  can be a negative numebr to substract
	fmt.Println("Please number to add/reduce stock (negative/positive int):")
	input, err = reader.ReadString('\n')
	if err != nil {
		// something went wrong return
		fmt.Println("Cannot read input. Please try again")
		return
	}

	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("Please enter a valid positive/negative integer.")
		return
	}

	// check valid number
	stockModifNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid positive/negative integer.")
		return
	}

	// check enough stock
	if existingProduct.Stock+stockModifNum < 0 {
		fmt.Printf("Not enough stock. Current stock: %s", existingProduct.Stock)
		return
	}

	// update stock
	existingProduct.Stock += stockModifNum
	productRepo.UpdateProduct(existingProduct)
}
