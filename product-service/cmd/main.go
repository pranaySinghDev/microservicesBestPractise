package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	database "github.com/pranaySinghDev/goSAK/database"
	"github.com/pranaySinghDev/goSAK/database/config"
)

// User struct
type Product struct {
	ID       string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func main() {
	app := fiber.New()
	//connect to db
	dbURL := os.Getenv("db_url")
	if dbURL == "" {
		log.Fatalf("Database connection string missing")
	}
	port := os.Getenv("product_port")
	if port == "" {
		log.Fatalf("TCP connection port string missing")
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
		msg := fmt.Sprintf("Hello, you have reached the product service 👋!")
		return c.SendString(msg)
	})

	//GetAll products
	app.Get("/products", func(c *fiber.Ctx) error {
		var Products []Product = make([]Product, 0)
		err := db.GetAll(c.Context(), "awesomeApp", "products", &Products)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(Products)
	})

	// GET a product
	app.Get("/products/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		product := &Product{}
		err := db.GetByID(c.Context(), "awesomeApp", "products", id, product)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(200).JSON(product)
	})

	// POST /products
	app.Post("/products", func(c *fiber.Ctx) error {
		// New Product struct
		product := new(Product)
		// Parse body into struct
		if err := c.BodyParser(product); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		product.ID = guuid.New().String()

		// insert the record
		err := db.Create(c.Context(), "awesomeApp", "products", product)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(201).JSON(product)
	})

	addr := fmt.Sprintf(":%s", port)
	log.Fatal(app.Listen(addr))
}
