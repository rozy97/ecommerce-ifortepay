package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rozy97/ecommerce-ifortepay/request"
)

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	payload := new(request.Register)
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	result, err := uh.uu.Register(c.Context(), payload)
	if err != nil {
		return err
	}

	c.Status(http.StatusOK).JSON(result)

	return nil
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	payload := new(request.Login)
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	// TODO: implement request validator

	resp, err := uh.uu.Login(c.Context(), payload)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(resp)
}
