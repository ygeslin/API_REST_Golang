package main

import (
	// "net/http"
	"dep/API_REST_Golang/configs"
	"dep/API_REST_Golang/routes"

	"github.com/gin-gonic/gin"
)

// TODO implements tests in Go instead of bash with kind of asserts like in Python
// TODO add IP and port in the .env file for flexibility
// TODO dockerize app and db with docker compose
// TODO https with nginx
// // replace MongoDB Atlas by standalone db
func main() {
	router := gin.Default()

	// connect to Databbase
	configs.ConnectDB()

	// set routes
	routes.SetRoutes(router)

	// run server
	router.Run("localhost:8080")
}
