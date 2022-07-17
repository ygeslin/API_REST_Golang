package controllers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	// "github.com/gin-gonic/gin/binding"

	"dep/API_REST_Golang/configs"
	"dep/API_REST_Golang/models"
	"encoding/json"
	"io/ioutil"
	"log"
)

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
    // fmt.Println(payload);
    return payload
}

// * POST /add/users
func CreateUser (c *gin.Context) {
	fmt.Print("CreateUser Function\n")
	// Read the raw data of the body
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Deserialize the json Data
	var DataSet []models.User
	err = json.Unmarshal(jsonData, &DataSet)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	// fmt.Println(DataSet)

	collection := configs.GetCollection(configs.DB, "users")

	for _, item := range DataSet {
		// test := collection.FindOne( c, item.Email)
		test2, err:= collection.Find( c, bson.M{"id": item.ID})
		// fmt.Println(test)
		fmt.Println(test2)

		res, err := collection.InsertOne(context.Background(), item)
		if err != nil {
			return
		}
		if res == nil{
			return
		}
	}

	// response, _ := json.Marshal(res)
	// c.Send(response)
	fmt.Println(collection)

}
// * POST /login
func Login (c *gin.Context) {
	fmt.Print("Login Function\n")
}

// * DELETE /delete/user/:id
func DeleteUser (c *gin.Context) {
	fmt.Print("DeleteUser Function\n")
}

// * GET /user/:id
func GetAUser (c *gin.Context) {
	fmt.Print("GetAUser Function\n")
}

// * GET /users/list
func GetUserList (c *gin.Context) {
	fmt.Print("GetUserList Function\n")
}

// * UPDATE /user/:id
func UpdateAUser (c *gin.Context) {
	fmt.Print("UpdateAUser Function\n")
}