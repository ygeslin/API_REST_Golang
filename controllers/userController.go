package controller

import (
	// "net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

// * POST /add/users
func CreateUser (c *gin.Context) {
	fmt.Print("CreateUser Fonction\n")
}
// * POST /login
func Login (c *gin.Context) {
	fmt.Print("Login Fonction\n")
}

// * DELETE /delete/user/:id
func DeleteUser (c *gin.Context) {
	fmt.Print("DeleteUser Fonction\n")
}

// * GET /user/:id
func GetAUser (c *gin.Context) {
	fmt.Print("GetAUser Fonction\n")
}

// * GET /users/list
func GetUserList (c *gin.Context) {
	fmt.Print("GetUserList Fonction\n")
}

// * UPDATE /user/:id
func UpdateAUser (c *gin.Context) {
	fmt.Print("UpdateAUser Fonction\n")
}