package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Email string `json:"email"`
}

func getall(c *gin.Context) {
	var users []Product

	db.Find(&users)

	c.JSON(200, gin.H{"users": users})
}

func getbyid(c *gin.Context) {
	var user Product

	id := c.Param("id")
	db.First(&user, id)

	c.JSON(200, gin.H{"user": user})
}

func createuser(c *gin.Context) {
	var newproduct Product

	if err := c.BindJSON(&newproduct); err != nil {
		return
	}
	db.Create(&Product{Name: newproduct.Name, Age: newproduct.Age, Email: newproduct.Email})

	c.IndentedJSON(200, newproduct)
}

func updateuser(c *gin.Context) {
	id := c.Param("id")

	// get the data from req body
	var body Product
	c.Bind(&body)

	// find the product that we are updating
	var user Product
	db.First(&user, id)

	db.Model(&user).Updates(&Product{Name: body.Name, Age: body.Age, Email: body.Email})

	c.JSON(200, gin.H{"user": user})
}

func deleteuser(c *gin.Context) {
	id := c.Param("id")

	db.Delete(&Product{}, id)

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
	db.AutoMigrate(&Product{})

	router := gin.Default()
	router.GET("/getallusers", getall)
	router.GET("/getuserbyid/:id", getbyid)
	router.POST("/createuser", createuser)
	router.PATCH("/updateuser:id", updateuser)
	router.DELETE("/deleteuser/:id", deleteuser)
	router.Run("0.0.0.0:8080")
}
