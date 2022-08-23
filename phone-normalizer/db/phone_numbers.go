package db

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

type Number struct {
	ID     int
	Number string
}

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Seed() error {
	numbers := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, number := range numbers {
		if _, err := insertPhone(db.db, number); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) FindAll() ([]Number, error) {
	rows, err := db.db.Query("SELECT id, value FROM phone_numbers")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []Number

	for rows.Next() {
		var n Number
		if err := rows.Scan(&n.ID, &n.Number); err != nil {
			return nil, err
		}

		result = append(result, n)
	}

	return result, nil
}

func (db *DB) FindNumber(number string) (*Number, error) {
	var n Number

	err := db.db.QueryRow("SELECT id, value FROM phone_numbers WHERE value=$1", number).Scan(&n.ID, &n.Number)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &n, nil
}

func (db *DB) Delete(phone Number) error {
	query := `DELETE FROM phone_numbers WHERE id=$1`

	_, err := db.db.Exec(query, phone.ID)

	return err
}

func (db *DB) Update(phone Number) error {
	query := `UPDATE phone_numbers SET value=$2 WHERE id=$1`

	_, err := db.db.Exec(query, phone.ID, phone.Number)

	return err
}

func Setup(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)

	if err != nil {
		return err
	}

	err = createPhoneNumbersTable(db)

	if err != nil {
		return err
	}

	return db.Close()
}

func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)

	if err != nil {
		return err
	}

	err = resetDB(db, dbName)

	if err != nil {
		return err
	}

	return db.Close()
}

func createPhoneNumbersTable(db *sql.DB) error {
	query := `
  CREATE TABLE IF NOT EXISTS phone_numbers (
    id SERIAL,
    value VARCHAR(255)
  )
  `

	_, err := db.Exec(query)

	return err
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", name))

	return err
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	var id int
	query := "INSERT INTO phone_numbers(value) VALUES($1) RETURNING id"
	err := db.QueryRow(query, phone).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", name))

	if err != nil {
		// in case database already there
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "object_in_use" {
			return nil
		}
		return err
	}

	return createDB(db, name)
}
