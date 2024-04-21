package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User is the model in the database
type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Age  uint8              `bson:"age" json:"age"`
}

var db *mongo.Collection

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my fiber rest api, the operations are available in the following endpoints \n /users\n /users/:id\n /create\n /update/:id\n /delete/:id")
}

func getall(c *fiber.Ctx) error {
	// Get all users
	results := []User{}
	get, err := db.Find(context.Background(), bson.D{})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer get.Close(context.Background())

	// decode to struct and return in json
	err = get.All(context.Background(), &results)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"users": results})
}

func getbyid(c *fiber.Ctx) error {
	// get the id
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// decode to struct and return in json
	result := User{}
	err = db.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&result)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"user": result})
}
func createuser(c *fiber.Ctx) error {
	// validate the body
	newuser := User{}
	if err := c.BodyParser(&newuser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}
	// insert into the database
	res, err := db.InsertOne(context.Background(), newuser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"result": res,
	})
}

func updateuser(c *fiber.Ctx) error {
	newuser := User{}
	if err := c.BodyParser(&newuser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}
	// get the id
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	// update the user
	result, err := db.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": newuser})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update book",
			"message": err.Error(),
		})
	}
	// return the book
	return c.Status(200).JSON(fiber.Map{
		"result": newuser,
		"err":    result,
	})
}
func deleteuser(c *fiber.Ctx) error {
	// get the id
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	//delete the user
	result, err := db.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete book",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})

}

func main() {
	//db connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:pass@mongo:27017"))
	if err != nil {
		panic("Couldn't connect to Mongo")
	}
	// defining the database and "table"
	db = client.Database("fiberdb").Collection("users")

	//ROUTES
	app := fiber.New()
	app.Get("/", welcome)
	app.Get("/users", getall)
	app.Get("/users/:id", getbyid)
	app.Post("/create", createuser)
	app.Put("/update/:id", updateuser)
	app.Delete("/delete/:id", deleteuser)
	app.Listen("0.0.0.0:8080")
}
