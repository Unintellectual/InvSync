package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	dsn := "user:!arestan22@tcp(localhost:3306)/inventory_db"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}
	fmt.Println("Connected to MySQL")
}

func getFoodItems(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, category, quantity, price, expiry_date FROM food_items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var foodItems []map[string]interface{}
	for rows.Next() {
		var id int
		var name, category string
		var quantity int
		var price float64
		var expiryDate *string
		err := rows.Scan(&id, &name, &category, &quantity, &price, &expiryDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		foodItem := gin.H{
			"id":          id,
			"name":        name,
			"category":    category,
			"quantity":    quantity,
			"price":       price,
			"expiry_date": expiryDate,
		}
		foodItems = append(foodItems, foodItem)
	}
	c.JSON(http.StatusOK, foodItems)
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()
	r.GET("/items", getFoodItems)

	r.Run(":8080")
}
