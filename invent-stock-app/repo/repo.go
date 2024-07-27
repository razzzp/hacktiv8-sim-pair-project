package repo

import (
	"database/sql"
)

type Product struct {
	Id    int
	Name  string
	Price float64
	Stock int
}

type Staff struct {
	Name     string
	Email    string
	Position string
}

type ProductRepo interface {
	AddProduct(product *Product) error
	UpdateProduct(product *Product) (int, error)
	GetProductByName(name string) (*Product, error)
}

type StaffRepo interface {
	AddStaff(staff *Staff) error
}

type productRepo struct {
	db *sql.DB
}

type staffRepo struct {
	db *sql.DB
}

func CreateProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func CreateStaffRepo(db *sql.DB) StaffRepo {
	return &staffRepo{
		db: db,
	}
}

// adds a product to DB
func (pr *productRepo) AddProduct(product *Product) error {
	// moved from main
	query := `INSERT INTO Products (Name, Price, Stock)
		VALUES (?, ?, ?);`
	_, err := pr.db.Exec(query, product.Name, product.Price, product.Stock)
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepo) UpdateProduct(product *Product) (int, error) {
	// TODO
	return 0, nil
}

func (pr *productRepo) GetProductByName(name string) (*Product, error) {
	// TODO
	return nil, nil
}

func (sr *staffRepo) AddStaff(staff *Staff) error {
	query := `INSERT INTO Staff (Name, Email, Position)
		VALUES (?, ?, ?);`
	_, err := sr.db.Exec(query, staff.Name, staff.Email, staff.Position)
	if err != nil {
		return err
	}
	return nil
}