package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://eventpass.rohitludbe.com, http://eventpass.rohitludbe.com", // Allow all origins, or specify comma-separated list
		AllowMethods:     "GET,POST,PUT,DELETE",                                                                      // Allowed HTTP methods
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",                                              // Allowed request headers
		ExposeHeaders:    "Authorization, X-Custom-Header",
		AllowCredentials: true,
	}))

	// repositories

	eventRepository := repositories.NewEventRepository(database.DB.Db)
	ticketRepository := repositories.NewTicketRepository(database.DB.Db)
	authRepository := repositories.NewAuthRepository(database.DB.Db)

	//service
	authService := services.NewAuthService(authRepository)

	// routing
	server := app.Group("/api/v1")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	eventGroup := server.Group("/event", middlewares.AuthProtected(database.DB.Db))
	ticketGroup := server.Group("/ticket", middlewares.AuthProtected(database.DB.Db))
	handlers.NewEventHandler(eventGroup, eventRepository)
	handlers.NewTicketHandler(ticketGroup, ticketRepository)
	
	app.Listen("0.0.0.0:3001")
}
