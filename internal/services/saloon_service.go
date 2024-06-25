package services

import (
	"database/sql"
	"errors"

	"github.com/MorgarAkt/MorPOS/internal/models"
)

type SalonService struct {
	db *sql.DB
}

func NewSalonService(db *sql.DB) *SalonService {
	return &SalonService{db: db}
}

func (s *SalonService) CreateTable() error {
	checkTableExistsSQL := `SELECT name FROM sqlite_master WHERE type='table' AND name='salons';`
	row := s.db.QueryRow(checkTableExistsSQL)
	var tableName string
	err := row.Scan(&tableName)
	if err == sql.ErrNoRows { // Table does not exist, create it
		createTableSQL := `CREATE TABLE salons (
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

func (s *SalonService) Insert(salon interface{}) error {
	sl, ok := salon.(models.Salon)
	if !ok {
		return errors.New("invalid salon type")
	}

	stmt, err := s.db.Prepare("INSERT INTO salons (id, name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sl.ID.String(), sl.Name)
	return err
}

func (s *SalonService) Update(salon interface{}) error {
	sl, ok := salon.(models.Salon)
	if !ok {
		return errors.New("invalid salon type")
	}

	stmt, err := s.db.Prepare("UPDATE salons SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sl.Name, sl.ID.String())
	return err
}

func (s *SalonService) Delete(salon interface{}) error {
	sl, ok := salon.(models.Salon)
	if !ok {
		return errors.New("invalid salon type")
	}

	stmt, err := s.db.Prepare("DELETE FROM salons WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sl.ID.String())
	return err
}
