package main

import (
	"realChakrawarti/rest-event/db"
	"realChakrawarti/rest-event/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDb()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":42069") // localhost:42069
}
