package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	database "github.com/pranaySinghDev/goSAK/database"
	"github.com/pranaySinghDev/goSAK/database/config"
)

// User struct
type User struct {
	ID      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string  `json:"name"`
	Product string  `json:"product"`
	Age     float64 `json:"age"`
}
type Product struct {
	ID       string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

type GetUserResponse struct {
	ID      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string  `json:"name"`
	Product Product `json:"product"`
	Age     float64 `json:"age"`
}

func main() {
	app := fiber.New()
	//connect to db
	dbURL := os.Getenv("db_url")
	if dbURL == "" {
		log.Fatalf("Database connection string missing")
	}
	port := os.Getenv("user_port")
	if port == "" {
		log.Fatalf("TCP connection port string missing")
	}
	productURL := os.Getenv("product_url")
	if productURL == "" {
		log.Fatalf("Product service URL is missing")
	}
	db, err := database.Build(&config.DBConfig{
		Type: config.Mongodb,
		URL:  dbURL,
	})

	if err != nil {
		log.Fatalf("Couldn't build database Factory: %v", err)
	}

	// GET / base
	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, you have reached the base 👋!")
		return c.SendString(msg)
	})

	//Get all users
	app.Get("/users", func(c *fiber.Ctx) error {
		var Users []User = make([]User, 0)
		err := db.GetAll(c.Context(), "awesomeApp", "users", &Users)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(200).JSON(Users)
	})

	// GET a user
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := &User{}
		err := db.GetByID(c.Context(), "awesomeApp", "users", id, user)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		product, err := getProduct(user.Product, productURL)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		res := GetUserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Product: *product,
			Age:     user.Age,
		}
		return c.Status(200).JSON(res)
	})

	// POST /user
	app.Post("/users", func(c *fiber.Ctx) error {
		// New user struct
		user := new(User)
		// Parse body into struct
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		user.ID = guuid.New().String()

		// insert the record
		err := db.Create(c.Context(), "awesomeApp", "users", user)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(201).JSON(user)
	})
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(app.Listen(addr))
}

func getProduct(id, productURL string) (*Product, error) {

	url := fmt.Sprintf("%s/products/%s", productURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &Product{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
