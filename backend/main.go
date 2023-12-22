package main

import (
	"flag"

	"github.com/AtharvaWaghchoure/goreserve/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API Server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandlerGetUsers)
	apiv1.Get("/user/:id", api.HandlerGetUser)
	app.Listen(*listenAddr)
}
