package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohit-ludbe/ticket-booking-server-v1/database"
	"github.com/rohit-ludbe/ticket-booking-server-v1/handlers"
	"github.com/rohit-ludbe/ticket-booking-server-v1/middlewares"
	"github.com/rohit-ludbe/ticket-booking-server-v1/repositories"
	"github.com/rohit-ludbe/ticket-booking-server-v1/services"
)

func main() {

	database.ConnectDb()

	app := fiber.New(fiber.Config{
		AppName:      "ticketBooking",
		ServerHeader: "Fiber",
	})

	// repositories

	eventRepository := repositories.NewEventRepository(database.DB.Db)
	ticketRepository := repositories.NewTicketRepository(database.DB.Db)
	authRepository := repositories.NewAuthRepository(database.DB.Db)

	//service
	authService := services.NewAuthService(authRepository)

	// routing
	server := app.Group("/api/v1")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(database.DB.Db))

	// handler

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(":3001")
}
