package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func welcome(c *gin.Context) {
	message := "Welcome to my gin rest api, the operations are available in the following endpoints /users /users/:id /create /update/:id /delete/:id"
	c.JSON(200, message)
}

func getall(c *gin.Context) {
	var users []User

	db.Find(&users)

	c.JSON(200, gin.H{"users": users})
}

func getbyid(c *gin.Context) {
	var user User

	id := c.Param("id")
	db.First(&user, id)

	c.JSON(200, gin.H{"user": user})
}

func createuser(c *gin.Context) {
	var newuser User

	if err := c.BindJSON(&newuser); err != nil {
		return 
	}
	db.Create(&User{Name: newuser.Name, Age: newuser.Age})

	c.IndentedJSON(200, newuser)
}

func updateuser(c *gin.Context) {
	id := c.Param("id")

	// get the data from req body
	var body User
	c.Bind(&body)

	// find the User that we are updating
	var user User
	db.First(&user, id)

	db.Model(&user).Updates(&User{Name: body.Name, Age: body.Age})

	c.JSON(200, gin.H{"user": user})
}

func deleteuser(c *gin.Context) {
	id := c.Param("id")

	db.Delete(&User{}, id)

	c.JSON(200, gin.H{"message": "user deleted"})
}

func main() {
	var err error
	dsn := "root:pass@tcp(gormdb:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema to mysql
	db.AutoMigrate(&User{})

	app := gin.Default()
	app.GET("/", welcome)
	app.GET("/users", getall)
	app.GET("/users/:id", getbyid)
	app.POST("/create", createuser)
	app.PATCH("/update/:id", updateuser)
	app.DELETE("/delete/:id", deleteuser)
	app.Run("0.0.0.0:8080")
}
