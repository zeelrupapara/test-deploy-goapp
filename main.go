package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func ConnectENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("ENV not Connected")
	}
	fmt.Println("env loaded")
}

func main() {
	ConnectENV()
	fmt.Println("not need env")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	fmt.Println("Listen app")
	err := app.Listen(os.ExpandEnv(`:${PORT}`))
	if err != nil {
		fmt.Println("Listen===>", err)
	}
}
