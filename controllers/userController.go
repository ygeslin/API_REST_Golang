package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	// "github.com/gin-gonic/gin/binding"

	"dep/API_REST_Golang/configs"
	"dep/API_REST_Golang/models"
	"dep/API_REST_Golang/utils"
	"encoding/json"
	"io/ioutil"
	"log"
)

// TODO change response error code to fit with RFC
// TODO Add middleware in all controllers or in the router ?
// TODO restore user ID from jwt, to limit access to only his profile

	var collection = configs.GetCollection(configs.DB, "users")

// TODO refactor the function to be more flexible, to handle both type (user and []user) in the same function (Maybe with typeof and if)
// TODO Add a validator for the data
func extractJsonArrayRequestBody (body io.ReadCloser) []models.User {
	// Read the raw data of the body
	jsonData, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Deserialize the json Data
	var DataSet []models.User
	err = json.Unmarshal(jsonData, &DataSet)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return DataSet
}
func extractJsonRequestBody (body io.ReadCloser) models.User {
	// Read the raw data of the body
	jsonData, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Deserialize the json Data
	var DataSet models.User
	err = json.Unmarshal(jsonData, &DataSet)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return DataSet
}

func extractJsonLoginBody (body io.ReadCloser) models.SignInInput {
	jsonData, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Deserialize the json Data
	var DataSet models.SignInInput
	err = json.Unmarshal(jsonData, &DataSet)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return DataSet
}

func createUserFile(fileName, data string) {
		filepath := "./userFiles/" + fileName
		// create or truncate the file
		f, err := os.Create(filepath)
		if err != nil {
			fmt.Println(err)
		}
		// close the file with defer
		defer f.Close()
		// write into the file
		f.WriteString(data);
		fmt.Println("Data file of the user was successfully created")

}

func deleteUserFile(fileName string) {
		filepath := "./userFiles/" + fileName
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
	}
}

func addUserRoutine (c *gin.Context, item models.User) {
		// query is the filter
		query := bson.M{"id": item.ID}
		// try to find if the id of the user already exist in the DB
		tmp := collection.FindOne(c, query)

		var user models.User
		// Decode user ID from bson
		tmp.Decode(&user)
		// Hash password and replace it in the user model
		hashPassword, err := utils.HashPassword(item.Password)
		if err != nil {
			return
		}
		item.Password = hashPassword
		// If the user isn't in the DB, add it
		if len(user.ID) == 0 {
		res, err := collection.InsertOne(context.Background(), item)
		if err != nil {
			return
		}
		fmt.Printf("--------------------------------------------------\n")
		fmt.Printf("InsertOne document with _id: %v \n", res.InsertedID)
		 // create the file
		createUserFile(item.ID, item.Data);

		// if User is already in the DB, just print info and don't add it
		} else {
		fmt.Printf("--------------------------------------------------\n")
		fmt.Println("The user with ID:'", user.ID, "' is already register in the Database")
		}
	return
}

// * POST /add/users
//// Add thread for each user creation
//// check if the user exist in db before create DONE
// TODO Add Real time arg if "registered" field is empty when creation user
func CreateUser (c *gin.Context) {
	fmt.Print("CreateUser Function\n")
	DataSet := extractJsonArrayRequestBody(c.Request.Body)

	for _, item := range DataSet {
		go addUserRoutine(c, item)
	}
	// return with json
}
// * POST /login
// TODO limit the size and complexity of username and password (use regex)
// TODO Add refresh token
// TODO automate test with curl
// TODO add pepper to protect against rainbow table
// TODO split login function into multiple functions
func Login (c *gin.Context) {
	fmt.Print("Login Function\n")
	credentials := extractJsonLoginBody(c.Request.Body)
	var user models.User
	err := collection.FindOne(c, bson.M{"id": credentials.ID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Sorry, user not found"})
		return
	}
	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}
	privateKey := configs.GetPrivateKey();
	// Generate Tokens
	access_token, err := utils.CreateToken(time.Duration(3600000000000), user.ID, privateKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(time.Duration(3600000000000), user.ID, privateKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, 36000000, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, 36000000, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

// * DELETE /delete/user/:id
func DeleteUser (c *gin.Context) {
	fmt.Print("DeleteUser Function\n")
		if IsAuthorized(c) == false{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
	userId := c.Param("id")
	var user models.User
	err := collection.FindOne(c, bson.M{"id": userId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusOK, "Sorry, user not found")
		return
	}
	res, err := collection.DeleteOne(c, bson.M{"id": userId})
	if err != nil {
    log.Fatal(err)
	}
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("DeleteOne removed %v document(s)\n", res.DeletedCount)
	deleteUserFile(userId)
}

func IsAuthorized(c *gin.Context) bool {
		// get token from Authorization header
		a:=  c.Request.Header["Authorization"]
		if (len(a)) == 0 {
			return false
		}
		token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		fmt.Print(token)
		res, error := utils.ValidateToken(token, configs.GetPublicKey())
		if res == nil || error != nil {
			return true
		}
		return false
}

// * GET /user/:id
func GetAUser (c *gin.Context) {
		fmt.Print("GetAUser Function\n")
		if IsAuthorized(c) == false{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		userId := c.Param("id")
		var user models.User
		err := collection.FindOne(c, bson.M{"id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "Sorry, user not found"})
			return
		}
		c.JSON(http.StatusOK, user)
}

// * GET /users/list
func GetUserList (c *gin.Context) {
	fmt.Print("GetUserList Function\n")
		if IsAuthorized(c) == false{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
		var userList []models.User
		res, err := collection.Find(c, bson.M{})
		if err != nil {
			c.JSON(http.StatusOK, "Sorry, no user in DataBase")
			return
		}
		for res.Next(c) {
			var singleUser models.User
			if err = res.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError,"Error while Decoding bson file from DataBase" )
			}

			userList = append(userList, singleUser)
		}
		c.JSON(http.StatusOK, userList)
}

// * UPDATE /user/:id
// TODO optimize the update with UpdateOne instead of ReplaceOne
// TODO implement more flexible input from the body (to be able to input only few fields in the json, whitout the User Struct)
func UpdateAUser (c *gin.Context) {
	fmt.Print("UpdateAUser Function\n")
		if IsAuthorized(c) == false{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
	userId := c.Param("id")
	var user models.User
	err := collection.FindOne(c, bson.M{"id": userId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusOK, "Sorry, user not found")
		return
	}
	requestBody := extractJsonRequestBody(c.Request.Body)
	if requestBody.ID != user.ID {
		fmt.Printf("--------------------------------------------------\n")
		fmt.Println("Error, user ID from URL and the request body are differents")
		return
	}
	if user.Data != requestBody.Data {
		createUserFile(userId, requestBody.Data)
	}
	res, err:= collection.ReplaceOne(c, bson.M{"id": userId}, requestBody)
	if err != nil {
    log.Fatal(err)
	}
		fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("Update %v document(s)\n", res.ModifiedCount)
}