package main

import (
	"fmt"

	_ "github.com/lib/pq"
	phoneDb "github.com/omarahm3/gogo/phone-normalizer/db"
	"github.com/omarahm3/gogo/phone-normalizer/lib"
)

const (
	host     = "localhost"
	user     = "root"
	password = "root"
	dbname   = "test"
	port     = 5432
	sslmode  = "disable"
)

func main() {
	connStr := fmt.Sprintf("host=%s password=%s user=%s port=%d sslmode=%s", host, password, user, port, sslmode)
  check(phoneDb.Reset("postgres", connStr, dbname))

	connStr = fmt.Sprintf("%s dbname=%s", connStr, dbname)
  check(phoneDb.Setup("postgres", connStr))

  db, err := phoneDb.Open("postgres", connStr)
	check(err)
	defer db.Close()

	check(db.Seed())

	numbers, err := db.FindAll()
	check(err)

	for _, n := range numbers {
		number := lib.Normalize(n.Number)

		if number == n.Number {
			fmt.Println("Nothing changed")
			continue
		}

		fmt.Println("Updating", n.Number)
		existing, err := db.FindNumber(number)
		check(err)

		if existing != nil {
			// delete number
      check(db.Delete(n))
		} else {
			// update number
			n.Number = number
			check(db.Update(n))
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
