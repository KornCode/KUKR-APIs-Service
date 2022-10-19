package handler

import (
	"github.com/KornCode/KUKR-APIs-Service/app/errs"
	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.RequestError:
		return c.Status(e.StatusCode).JSON(fiber.Map{
			"message": e.Message,
		})
	default:
		return err
	}
}
