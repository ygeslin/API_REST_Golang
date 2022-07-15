
// * POST /add/users
func CreateUser (c *gin.Context) {
  c.JSON(http.StatusCreated, response)
}
// * POST /login
func Login (c *gin.Context) {
  c.JSON(http.StatusCreated, response)
}

// * DELETE /delete/user/:id
func DeleteUser (c *gin.Context) {
  c.JSON(http.StatusCreated, response)
}

// * GET /user/:id
func getAUser (c *gin.Context) {
  c.JSON(http.StatusCreated, response)
}

// * GET /users/list
func getUserList (c *gin.Context) {
  c.JSON(http.StatusCreated, response)
}

// * UPDATE /user/:id
func UpdateAUser (c *gin.Context) {
  c.JSON(http.StatusCreated, response)
}