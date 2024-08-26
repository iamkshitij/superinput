package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type RequestBody struct {
	RequestInput int `json:"requestInput"`
}

type ResponseBody struct {
	ResponseOutput int `json:"result"`
}

func main() {
	app := fiber.New()

	app.Post("/input", func(c *fiber.Ctx) error {
		log.Infof("Integer from Request: %s", c.Body())

		p := new(RequestBody)

		if err := c.BodyParser(p); err != nil {

			return err
		}

		log.Infof("Value from request", p.RequestInput) //

		res := new(ResponseBody)
		res.ResponseOutput = p.RequestInput

		return c.Status(200).JSON(res)
	})

	app.Listen(":3000")
}
