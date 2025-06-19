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
	"os"

	_ "github.com/AshalIbrahim/ginApi/docs" // 👈 correct path to docs
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

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

	if err := DB.AutoMigrate(&Users{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
}

func main() {
	initDB()
	r := gin.Default()

	// Grouping routes under /api
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users", AllUsers)
			v1.POST("/users", createUser)
			v1.PUT("/users/:id", updateUser)
			v1.DELETE("/users/:id", deleteUser) 
		}
		v2 := api.Group("/v2")
		{
			v2.GET("/", V2api)
		}
	}

	// Serve custom Swagger UI and Swagger docs
	r.StaticFile("/swagger-ui", "./swagger-ui.html")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
