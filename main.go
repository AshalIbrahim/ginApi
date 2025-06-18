package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// User model with GORM tags
type Users struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var DB *gorm.DB

func initDB() {
	
	dsn := "host=localhost user=root password=admin dbname=practicedb port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate will create or update the table structure automatically
	err = DB.AutoMigrate(&Users{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}

func main() {
	initDB()
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		var users []Users
		result := DB.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

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

	r.Run(":8080")
}
// This code sets up a simple REST API using Gin and GORM with PostgreSQL.
// It includes endpoints to retrieve all users and create a new user.
// Make sure to have PostgreSQL running and the database 'practiceDB' created before running this code.