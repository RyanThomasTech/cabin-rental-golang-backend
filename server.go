package main

import (
	//"database/sql"
	//"fmt"
	//"log"
	//"os"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

/*type Customer struct {
	firstName string
	lastName  string
}*/

type RentalRecord struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"date"`
	RentedTo string    `json:"rentedTo"`
	RentedBy string    `json:"rentedBy"`
}

var rentals = []RentalRecord{
	{ID: 0, Date: time.Date(2021, time.November, 21, 12, 0, 0, 0, time.UTC), RentedTo: "Ryan", RentedBy: "Ryan"},
	{ID: 1, Date: time.Date(2021, time.November, 22, 12, 0, 0, 0, time.UTC), RentedTo: "Ryan", RentedBy: "Ryan"},
	{ID: 2, Date: time.Date(2021, time.November, 27, 12, 0, 0, 0, time.UTC), RentedTo: "Toni", RentedBy: "Ryan"},
}

func main() {
	//load the environment variables
	//err := godotenv.Load()

	//router using custom middleware for CORS purposes
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(CORSMiddleware())
	r.GET("/rentals", getRentals)

	r.Run("localhost:8080")

	/*
		router := gin.Default()
		router.GET("/rentals", getRentals)
		router.GET("/rentals/:id", getRental)
		router.POST("/rentals", postRental)

		router.Run("localhost:8080")
	*/

	/*if err != nil {
		log.Fatal("error loading env file")
	}
	jamies, err := getJamies()
	if err != nil {
		log.Fatal("getjamies failed")
	}
	for _, cus := range jamies {
		fmt.Printf("Customer: %s %s", cus.firstName, cus.lastName)
	}*/
}

//copied from StackOverflow for educational purposes, can use official gin CORS methods at github.com/gin-contrib/cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //TODO: change * to specific URL, otherwise it's public
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func getRentals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rentals)
}

func getRental(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "non-int id found"})
		return
	}

	for _, a := range rentals {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rental not found"})
}

func postRental(c *gin.Context) {
	var newRecord RentalRecord

	if err := c.BindJSON(&newRecord); err != nil {
		return
	}

	rentals = append(rentals, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}

/*func getJamies() ([]Customer, error) {
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
*/
