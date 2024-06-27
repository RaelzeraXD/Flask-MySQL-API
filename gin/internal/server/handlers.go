package server

import (
	"github.com/RaelzeraXD/api/gin/internal/database"
	"github.com/RaelzeraXD/api/gin/internal/models"
	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	message := "Welcome to my gin rest api, the operations are available in the following endpoints /users /users/:id /create /update/:id /delete/:id"
	c.JSON(200, message)
}

func Getall(c *gin.Context) {
	var users []models.User

	database.Connect().Find(&users)

	c.JSON(200, gin.H{"users": users})
}

func Getbyid(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	database.Connect().First(&user, id)

	c.JSON(200, gin.H{"user": user})
}

func Createuser(c *gin.Context) {
	var newuser models.User

	if err := c.BindJSON(&newuser); err != nil {
		return
	}
	database.Connect().Create(&models.User{Name: newuser.Name, Age: newuser.Age})

	c.IndentedJSON(200, newuser)
}

func Updateuser(c *gin.Context) {
	id := c.Param("id")

	// get the data from req body
	var body models.User
	c.Bind(&body)

	// find the User that we are updating
	var user models.User
	database.Connect().First(&user, id)

	database.Connect().Model(&user).Updates(&models.User{Name: body.Name, Age: body.Age})

	c.JSON(200, gin.H{"user": user})
}

func Deleteuser(c *gin.Context) {
	id := c.Param("id")

	database.Connect().Delete(&models.User{}, id)

	c.JSON(200, gin.H{"message": "user deleted"})
}
