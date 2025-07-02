package main

import (
	"github.com/dhofa/gofiber-clean-arch/config"
	"github.com/dhofa/gofiber-clean-arch/infrastructure/db"
	"github.com/dhofa/gofiber-clean-arch/infrastructure/router"
	"github.com/dhofa/gofiber-clean-arch/internal/handler"
	"github.com/dhofa/gofiber-clean-arch/internal/repository"
	"github.com/dhofa/gofiber-clean-arch/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}

	// Users setup
	userRepo := repository.NewUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	app := fiber.New()

	// Inject all handlers ke router registry
	routes := &router.RouteRegistry{
		UserHandler: userHandler,
	}

	router.Setup(app, routes)
	app.Listen(":3000")
}
