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

type TestJson struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Number string `json:"number"`
}

func main() {
	ConnectENV()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/api/v1/test", func(c *fiber.Ctx) error {
		test := &TestJson{
			Name:   "jeel",
			Age:    "21",
			Number: "0",
		}
		return c.JSON(test)
	})
	app.Static("/favicon.ico", "./public/image/favicon.ico")

	err := app.Listen(os.ExpandEnv(`:${PORT}`))
	if err != nil {
		fmt.Println(err)
	}
}
