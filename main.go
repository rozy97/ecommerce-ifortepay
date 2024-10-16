package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rozy97/ecommerce-ifortepay/config"
	"github.com/rozy97/ecommerce-ifortepay/handler"
	"github.com/rozy97/ecommerce-ifortepay/lib"
	"github.com/rozy97/ecommerce-ifortepay/repository"
	"github.com/rozy97/ecommerce-ifortepay/usecase"
)

func main() {
	env := config.InitEnvironment()
	db, err := sqlx.Open("postgres", env.PostgresURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	clock := lib.NewClock()
	userUsecase := usecase.NewUserUsecase(userRepository, clock, env)
	userHandler := handler.NewUserHandler(userUsecase)
	// middleware := middleware.NewMiddleware(env)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Welcome to ifortepay ecommerce API"})
	})

	userRoutes := app.Group("/user")
	userRoutes.Post("/register", userHandler.Register)
	userRoutes.Post("/login", userHandler.Login)
	err = app.Listen(fmt.Sprintf(":%s", env.AppPort))
	if err != nil {
		log.Fatalln(err)
	}
}
