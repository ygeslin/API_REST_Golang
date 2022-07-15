package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "dep/API_REST_Golang/configs"

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
var users = user {
    Id: "1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd",
    Password: "CGUsfQkS06lo7LM2guSV",
    IsActive: false,
    Balance: "$2,547.50",
    Age: 27,
    Name: "Nikki Farley",
    Gender: "female",
    Company: "ANIVET",
    Email: "nikkifarley@anivet.com",
    Phone: "+1 (868) 439-2675",
    Address: "588 Schaefer Street, Falconaire, Missouri, 9457",
    About: "Commodo minim fugiat est fugiat sunt duis consectetur fugiat Lorem sunt. Dolor nulla nisi quis deserunt occaecat tempor reprehenderit. Amet sunt incididunt cillum magna id. Mollit reprehenderit irure non sunt aute consectetur adipisicing. Pariatur veniam proident ut nulla culpa eiusmod ipsum occaecat.\r\n",
    Registered: "2014-10-12T04:14:20 -02:00",
    Latitude: -10.647121,
    Longitude: 126.04006,
    Data:"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",
}

func getUser(c *gin.Context) {
       c.IndentedJSON(http.StatusOK, users)
}

func main() {
    // fmt.Println("Hello, World!")
    // fmt.Println(users)
    // init gin instance
    router := gin.Default()

    // connect to Databbase
    configs.ConnectDB()

    // set routes

    // run server
    router.GET("/user", getUser)
          router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "data": "Hello from Gin-gonic & mongoDB",
                })
        })
    router.Run("localhost:8080")
}
