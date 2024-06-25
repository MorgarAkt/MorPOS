package services

import (
	"database/sql"
	"errors"

	"github.com/MorgarAkt/MorPOS/internal/models"
)

type TableService struct {
	db *sql.DB
}

func NewTableService(db *sql.DB) *TableService {
	return &TableService{db: db}
}

func (s *TableService) CreateTable() error {
	checkTableExistsSQL := `SELECT name FROM sqlite_master WHERE type='table' AND name='tables';`
	row := s.db.QueryRow(checkTableExistsSQL)
	var tableName string
	err := row.Scan(&tableName)
	if err == sql.ErrNoRows { // Table does not exist, create it
		createTableSQL := `CREATE TABLE tables (
			id TEXT PRIMARY KEY,
			number INTEGER NOT NULL,
			totalBill REAL NOT NULL
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

func (s *TableService) Insert(table interface{}) error {
	t, ok := table.(models.Table)
	if !ok {
		return errors.New("invalid table type")
	}

	stmt, err := s.db.Prepare("INSERT INTO salon_tables (id, salon_id, number, totalBill) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.ID.String(), t.SalonID.String(), t.Number, t.TotalBill)
	return err
}

func (s *TableService) Update(table interface{}) error {
	t, ok := table.(models.Table)
	if !ok {
		return errors.New("invalid table type")
	}

	stmt, err := s.db.Prepare("UPDATE salon_tables SET salon_id = ?, number = ?, totalBill = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.SalonID.String(), t.Number, t.TotalBill, t.ID.String())
	return err
}

func (s *TableService) Delete(table interface{}) error {
	t, ok := table.(models.Table)
	if !ok {
		return errors.New("invalid table type")
	}

	stmt, err := s.db.Prepare("DELETE FROM salon_tables WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.ID.String())
	return err
}
