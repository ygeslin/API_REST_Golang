package main

import (
	// "net/http"
	"dep/API_REST_Golang/configs"
	"dep/API_REST_Golang/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type friend struct{
    Id          int        `json:"id"`
    Name        string     `json:"name"`
}
type user struct {
    Id          string     `json:"id"`
    Password    string     `json:"password"`
    IsActive    bool       `json:"isActive"`
    Balance     string     `json:"balance"`
    Age         int        `json:"age"`
    Name        string     `json:"name"`
    Gender      string     `json:"gender"`
    Company     string     `json:"company"`
    Email       string     `json:"email"`
    Phone       string     `json:"phone"`
    Address     string     `json:"address"`
    About       string     `json:"about"`
    Registered  string     `json:"registered"`
    Latitude    float32    `json:"Latitude"`
    Longitude   float32    `json:"Longitude"`
    Tags        []string   `json:"tags"`
    Friends     []friend   `json:"friends"`
    Data        string     `json:"data"`
}

func getUser(c *gin.Context) {
    //    c.IndentedJSON(http.StatusOK)
}

func importDataSet() []user {
    content, err := ioutil.ReadFile("./data/DataSet.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }

    // Now let's unmarshall the data into `payload`
    var payload []user
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
