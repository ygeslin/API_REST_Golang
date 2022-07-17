package main

import (
	// "net/http"
	"dep/API_REST_Golang/routes"
	"dep/API_REST_Golang/configs"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // connect to Databbase
    configs.ConnectDB()

    // set routes
    routes.SetRoutes(router)

    // run server
    // router.GET("/user", getUser)
    //       router.GET("/", func(c *gin.Context) {
    //             c.JSON(200, gin.H{
    //                     "data": "Hello from Gin-gonic & mongoDB",
    //             })
    //     })
    router.Run("localhost:8080")
}
