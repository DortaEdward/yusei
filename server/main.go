package main

import (
	"log"
	"github.com/DortaEdward/tcgStratServer/api"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	server := fiber.New();

	// Middleware
	server.Use(helmet.New())
	server.Use(cors.New())
	server.Use(limiter.New(limiter.Config{
		Max: 100,
	}))
	server.Use(logger.New())

	// Api Router
	api := api.NewApi(server)
	api.Run()


	_, clerkErr := clerk.NewClient("")
	if clerkErr != nil {
		log.Fatal("ERROR: Cannot create Clerk client\n",clerkErr)
	} 

	// Server Listen
	err := server.Listen(":5042"); if err != nil {
		log.Fatal("ERROR: There was an error creating the server \n",err)
	}
}

