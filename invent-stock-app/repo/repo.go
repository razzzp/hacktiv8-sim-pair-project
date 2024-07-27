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

type ProductRepo interface {
	AddProduct(product *Product) error
	UpdateProduct(product *Product) (int, error)
	GetProductByName(name string) (*Product, error)
}

type productRepo struct {
	db *sql.DB
}

func CreateProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
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
