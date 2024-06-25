package services

import (
	"database/sql"
	"errors"

	"github.com/MorgarAkt/MorPOS/internal/models"
)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) CreateTable() error {
	checkTableExistsSQL := `SELECT name FROM sqlite_master WHERE type='table' AND name='products';`
	row := s.db.QueryRow(checkTableExistsSQL)
	var tableName string
	err := row.Scan(&tableName)
	if err == sql.ErrNoRows { // Table does not exist, create it
		createTableSQL := `CREATE TABLE products (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			image BLOB,
			price REAL NOT NULL
		);`

		stmt, err := s.db.Prepare(createTableSQL)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	} else if err != nil { // Other error occurred
		return err
	}

	return nil
}

func (s *ProductService) Insert(product interface{}) error {
	p, ok := product.(models.Product)
	if !ok {
		return errors.New("invalid product type")
	}

	stmt, err := s.db.Prepare("INSERT INTO products (id, name, image, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID.String(), p.Name, p.Image, p.Price)
	return err
}

func (s *ProductService) Update(product interface{}) error {
	p, ok := product.(models.Product)
	if !ok {
		return errors.New("invalid product type")
	}

	stmt, err := s.db.Prepare("UPDATE products SET name = ?, image = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Image, p.Price, p.ID.String())
	return err
}

func (s *ProductService) Delete(product interface{}) error {
	p, ok := product.(models.Product)
	if !ok {
		return errors.New("invalid product type")
	}

	stmt, err := s.db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID.String())
	return err
}
