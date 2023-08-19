package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string
	Description string
}

func main() {
	dsn := `host=localhost user=postgres password=password dbname=postgres port=5432 `
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Post{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
