package main

import (
	// "net/http"
	"dep/API_REST_Golang/configs"
	"dep/API_REST_Golang/models"
	"dep/API_REST_Golang/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
    //    c.IndentedJSON(http.StatusOK)
}

func importDataSet() []models.User {
    content, err := ioutil.ReadFile("./data/DataSet.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }

    // Now let's unmarshall the data into `payload`
    var payload []models.User
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

    // Let's print the unmarshalled data!
    fmt.Println(payload);
    return payload
}

func main() {
    // fmt.Println("Hello, World!")
    // fmt.Println(users)
    // convert json data set into struct
    DataSet := importDataSet();
    fmt.Println(DataSet)
    // init gin instance
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
