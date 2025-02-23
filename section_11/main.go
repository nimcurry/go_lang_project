package main

import (
	"fmt"

	"apilibrary.com/db"
	"apilibrary.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
