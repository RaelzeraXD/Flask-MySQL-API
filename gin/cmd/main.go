package main

import (
	"log"

	"github.com/RaelzeraXD/api/gin/internal/models"
	"github.com/RaelzeraXD/api/gin/internal/server"
	"github.com/RaelzeraXD/api/gin/internal/database"
	"github.com/gin-gonic/gin"

)

func main() {
	// Migrate the schema to mysql
	if err := database.Connect().AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	r := gin.Default()
	r.GET("/", server.Welcome)
	r.GET("/users", server.Getall)
	r.GET("/users/:id", server.Getbyid)
	r.POST("/create", server.Createuser)
	r.PUT("/update/:id", server.Updateuser)
	r.DELETE("/delete/:id", server.Deleteuser)
	r.Run("0.0.0.0:8080")
}
