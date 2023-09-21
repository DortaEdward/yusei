package api

import ("github.com/gofiber/fiber/v2")

type Api struct{
	server *fiber.App
}

func NewApi(server *fiber.App) *Api{
	return &Api{
		server: server,
	}
}

func HandleHealthCheck(c *fiber.Ctx) error{
	data := map[string]interface{}{
		"message":"Health Check",
		"status": fiber.StatusOK,
	}
	c.Status(fiber.StatusOK)
	c.JSON(data)
	return nil
}

func (s *Api) Run(){
	routes := s.server.Group("/api")
	routes.Get("/",HandleHealthCheck)
}
