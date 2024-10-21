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

	err = db.Ping()
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
	v1 := app.Group("/v1")

	v1.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Welcome to ifortepay ecommerce API"})
	})

	v1.Post("/register", userHandler.Register)
	v1.Post("/login", userHandler.Login)

	v1.Get("/product", func(c *fiber.Ctx) error { panic("implement me") })             // get list products
	v1.Get("/product/:product_id", func(c *fiber.Ctx) error { panic("implement me") }) // get product detail

	v1.Post("/wishlist", func(c *fiber.Ctx) error { panic("implement me") }) // add product to wishlist
	v1.Get("/wishlist", func(c *fiber.Ctx) error { panic("implement me") })  // get list wishlist

	v1.Post("/cart", func(c *fiber.Ctx) error { panic("implement me") }) // add product to cart
	v1.Get("/cart", func(c *fiber.Ctx) error { panic("implement me") })  // get list cart item

	v1.Post("/order", func(c *fiber.Ctx) error { panic("implement me") })          // create order
	v1.Get("/order", func(c *fiber.Ctx) error { panic("implement me") })           // get list order
	v1.Get("/order/:order_id", func(c *fiber.Ctx) error { panic("implement me") }) // get order detail

	err = app.Listen(fmt.Sprintf(":%s", env.AppPort))
	if err != nil {
		log.Fatalln(err)
	}
}
