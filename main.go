package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/twincats/fiber-app/server/config"
	"github.com/twincats/fiber-app/server/database"
	"github.com/twincats/fiber-app/server/middleware"
	"github.com/twincats/fiber-app/server/router"
)

var (
	port = flag.String("port", ":4444", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	//load config
	config.LoadConfig()

	// Create View Engine
	engine := html.New("./server/views", ".html")

	// Create fiber app
	app := fiber.New(fiber.Config{
		Views:   engine,
		Prefork: *prod,
	})

	// Database setup
	database.InitDatabase()

	// Router Setup
	router.SetupRouter(app)

	// Static File Setup
	app.Static("/", "./public")

	// Middleware
	app.Use(middleware.NotFound())

	// Listen on Port 4444
	log.Fatal(app.Listen(*port))
}
