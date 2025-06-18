package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv" 
)

//a:=[5]int{}

type Users struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

	var list=[]Users{
		{Id: 1,Name: "Alice", Age:30},
	}

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context){
		c.JSON(http.StatusOK, list)
	})

	r.POST("/users", func(c *gin.Context){

		//fmt.Println("User details:",c.body)

		var newUser Users
		if err := c.ShouldBindJSON(&newUser); err !=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
			return
		}
		newUser.Id=len(list) + 1
		list = append(list,newUser)

		c.JSON(http.StatusCreated, newUser)
	})

	r.Run(":8080")
}
