package controllers

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	// "github.com/gin-gonic/gin/binding"

	"dep/API_REST_Golang/configs"
	"dep/API_REST_Golang/models"
	"encoding/json"
	"io/ioutil"
	"log"
)

// func importDataSet() []models.User {
//     content, err := ioutil.ReadFile("./data/DataSet.json")
//     if err != nil {
//         log.Fatal("Error when opening file: ", err)
//     }

//     // Now let's unmarshall the data into `payload`
//     var payload []models.User
//     err = json.Unmarshal(content, &payload)
//     if err != nil {
//         log.Fatal("Error during Unmarshal(): ", err)
//     }

//     // Let's print the unmarshalled data!
//     // fmt.Println(payload);
//     return payload
// }

// * POST /add/users
// TODO Add thread for each user creation
//// check if the user exist in db before create DONE
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
		// query is the filter
		query := bson.M{"id": item.ID}
		// try to find if the id of the user already exist in the DB
		tmp := collection.FindOne(c, query)

		var user models.User
		// Decode user ID from bson
		tmp.Decode(&user)
		// If the user isn't in the DB, add it
		if len(user.ID) == 0 {
		res, err := collection.InsertOne(context.Background(), item)
		if err != nil {
			return
		}
		if res.InsertedID != nil{
		fmt.Println("The user with ID:'", item.ID, "' was successfully added")
		}
		 // create the file
		 filepath := "./userFiles/" + item.ID
		f, err := os.Create(filepath)
		if err != nil {
			fmt.Println(err)
		}
		// close the file with defer
		defer f.Close()
		f.WriteString(item.Data);
		fmt.Println("Data file of the user was successfully created")

		// if User is already in the DB, just print info and don't add it
		} else {
		fmt.Println("The user with ID:'", user.ID, "' is already register in the Database")
		}
	}

	// response, _ := json.Marshal(res)
	// c.Send(response)
	// fmt.Println(collection)

}
// * POST /login
// TODO Add jwt Auth
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
	// collection := configs.GetCollection(configs.DB, "users")
}

// * GET /users/list
func GetUserList (c *gin.Context) {
	fmt.Print("GetUserList Function\n")
}

// * UPDATE /user/:id
func UpdateAUser (c *gin.Context) {
	fmt.Print("UpdateAUser Function\n")
}