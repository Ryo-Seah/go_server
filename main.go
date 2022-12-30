package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Entry struct {
	ID       int    `json:"ID"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Likes    int    `json:"likes"`
	Body     string `json:"body"`
	Owner    string `json:"owner"`
}

type Auth struct {
	Username string `json:"username"`
	Passwors string `json:"passwors"`
}

func main() {
	fmt.Print("Hello world.......HHHH")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	entries := []Entry{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("IT WOrks niggas")
	})

	app.Post("/api/entry", func(c *fiber.Ctx) error {
		entry := &Entry{}

		if err := c.BodyParser(entry); err != nil {
			return err
		}
		entry.ID = len(entries) + 1
		entries = append(entries, *entry)
		return c.JSON(entries)
	})

	app.Patch("/api/entry/:id/like", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range entries {
			if t.ID == id {
				entries[i].Likes += 1
				break
			}
		}
		return c.JSON(entries)
	})

	app.Get("/api/entry", func(c *fiber.Ctx) error {
		return c.JSON(entries)
	})

	log.Fatal(app.Listen(":3002"))
}
