// @title           Gin API Example
// @version         1.0
// @description     This is a sample server built with Gin and GORM.
// @host            localhost:8080
// @BasePath        /
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	_ "github.com/AshalIbrahim/ginApi/docs" // 👈 replace with actual module name
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)


// User model with GORM tags
type Users struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var DB *gorm.DB

func initDB() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get variables from environment
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&Users{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
// 
func main() {
	initDB()
	r := gin.Default()

// @Summary      Get all users
// @Description  Returns a list of all users
// @Tags         users
// @Produce      json
// @Success      200  {array}  Users
// @Router       /users [get]
r.GET("/users", func(c *gin.Context) {
	var users []Users
	result := DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
})

// @Summary      Create a user
// @Description  Adds a new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  Users  true  "User to create"
// @Success      201   {object}  Users
// @Router       /users [post]
r.POST("/users", func(c *gin.Context) {
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
})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
// This code sets up a simple REST API using Gin and GORM with PostgreSQL.
// It includes endpoints to retrieve all users and create a new user.
// Make sure to have PostgreSQL running and the database 'practiceDB' created before running this code.