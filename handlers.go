package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)


// allUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all users
// @Tags         users
// @Produce      json
// @Success      200  {array}  Users
// @Router       /users [get]
func AllUsers(c *gin.Context) {
	var users []Users
	result := DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}


// createUser godoc
// @Summary      Create a user
// @Description  Adds a new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  Users  true  "User to create"
// @Success      201   {object}  Users
// @Router       /users [post]
func createUser(c *gin.Context) {
	var newUser Users
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}