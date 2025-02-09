package main

import (
	"log"

	"github.com/Unintellectual/database"
	"github.com/Unintellectual/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()
	defer database.CloseDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
