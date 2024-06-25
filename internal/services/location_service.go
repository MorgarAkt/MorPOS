package services

import (
	"database/sql"
	"errors"

	"github.com/MorgarAkt/MorPOS/internal/models"
)

type LocationService struct {
	db *sql.DB
}

func NewLocationService(db *sql.DB) *LocationService {
	return &LocationService{db: db}
}

func (s *LocationService) CreateTable() error {
	checkTableExistsSQL := `SELECT name FROM sqlite_master WHERE type='table' AND name='locations';`
	row := s.db.QueryRow(checkTableExistsSQL)
	var tableName string
	err := row.Scan(&tableName)
	if err == sql.ErrNoRows {
		createTableSQL := `CREATE TABLE locations (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL
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

func (s *LocationService) Insert(location interface{}) error {
	l, ok := location.(models.Location)
	if !ok {
		return errors.New("invalid location type")
	}

	stmt, err := s.db.Prepare("INSERT INTO locations (id, name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(l.ID.String(), l.Name)
	return err
}

func (s *LocationService) Update(location interface{}) error {
	l, ok := location.(models.Location)
	if !ok {
		return errors.New("invalid location type")
	}

	stmt, err := s.db.Prepare("UPDATE locations SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(l.Name, l.ID.String())
	return err
}

func (s *LocationService) Delete(location interface{}) error {
	l, ok := location.(models.Location)
	if !ok {
		return errors.New("invalid location type")
	}

	stmt, err := s.db.Prepare("DELETE FROM locations WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(l.ID.String())
	return err
}
