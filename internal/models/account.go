package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Account struct {
	ID      int32
	Name    string
	Balance float64
	Created time.Time
}

type AccountModel struct {
	DB *sql.DB
}

func (m *AccountModel) Insert(name string, balance float64) (int, error) {
	var newID int
	stmt := `
		INSERT INTO Account (name, balance)
		OUTPUT Inserted.ID
		VALUES (?, ?)`

	err := m.DB.QueryRow(stmt, name, balance).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *AccountModel) Get(id int) (Account, error) {
	stmt := `SELECT ID, name, balance, created FROM Account WHERE ID = ?`

	row := m.DB.QueryRow(stmt, id)
	var s Account
	err := row.Scan(&s.ID, &s.Name, &s.Balance, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Account{}, ErrNoRecord
		} else {
			return Account{}, err
		}
	}

	return s, nil
}

func (m *AccountModel) Latest() ([]Account, error) {
	rows, err := m.DB.Query(`
	SELECT TOP 10 ID, name, balance, created FROM Account
	ORDER BY created DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var s Account
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Balance,
			&s.Created,
		); err != nil {
			log.Println("Could not add value", err)
			continue
		}
		accounts = append(accounts, s)
	}

	return accounts, nil
}
