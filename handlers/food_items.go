package handlers

import (
	"net/http"

	"github.com/Unintellectual/InvSync/database"

	"github.com/gin-gonic/gin"
)

func GetFoodItems(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, category, quantity, price, expiry_date FROM food_items")
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
