package services

import (
	"database/sql"
	"errors"

	"github.com/MorgarAkt/MorPOS/internal/models"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateTable() error {
	checkTableExistsSQL := `SELECT name FROM sqlite_master WHERE type='table' AND name='users';`
	row := s.db.QueryRow(checkTableExistsSQL)
	var tableName string
	err := row.Scan(&tableName)
	if err == sql.ErrNoRows { // Table does not exist, create it
		createTableSQL := `CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL
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

func (s *UserService) Insert(user interface{}) error {
	u, ok := user.(models.User)
	if !ok {
		return errors.New("invalid user type")
	}

	stmt, err := s.db.Prepare("INSERT INTO users (id, fullName, username, password, role) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID.String(), u.FullName, u.Username, u.Password, u.Role)
	return err
}

func (s *UserService) Update(user interface{}) error {
	u, ok := user.(models.User)
	if !ok {
		return errors.New("invalid user type")
	}

	stmt, err := s.db.Prepare("UPDATE users SET fullName = ?, username = ?, password = ?, role = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FullName, u.Username, u.Password, u.Role, u.ID.String())
	return err
}

func (s *UserService) Delete(user interface{}) error {
	u, ok := user.(models.User)
	if !ok {
		return errors.New("invalid user type")
	}

	stmt, err := s.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID.String())
	return err
}
