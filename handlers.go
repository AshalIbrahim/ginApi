package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)


// allUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all users
// @Tags         api|users
// @Produce      json
// @Success      200  {array}  Users
// @Router       /api/v1/users [get]
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
// @Tags         api|users
// @Accept       json
// @Produce      json
// @Param        user  body  Users  true  "User to create"
// @Success      201   {object}  Users
// @Router       /api/v1/users [post]
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


// V2api godoc
// @Summary      V2 API Example
// @Description  This is an example endpoint for v2
// @Tags         api|users.V2
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/v2/ [get]
func V2api(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the V2 API endpoint"})
}