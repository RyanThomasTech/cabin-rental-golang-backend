package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Customer struct {
	firstName string
	lastName  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}
	jamies, err := getJamies()
	if err != nil {
		log.Fatal("getjamies failed")
	}
	for _, cus := range jamies {
		fmt.Printf("Customer: %s %s", cus.firstName, cus.lastName)
	}
}

func getJamies() ([]Customer, error) {
	var customers []Customer
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s", os.Getenv("DBNAME"), os.Getenv("DBUSER"), os.Getenv("DBPASS"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT first_name,last_name FROM customer WHERE first_name='Jamie'")
	if err != nil {
		return nil, fmt.Errorf("error on initial query")
	}
	defer rows.Close()
	for rows.Next() {
		var cus Customer
		if err := rows.Scan(&cus.firstName, &cus.lastName); err != nil {
			return nil, fmt.Errorf("failed during scan")
		}
		customers = append(customers, cus)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("the query failed")
	}
	return customers, nil
}
