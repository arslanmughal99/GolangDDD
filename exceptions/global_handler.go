package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"massivleads/logger"
)

// GlobalExceptionHandler handle any global exception
func GlobalExceptionHandler(ctx *fiber.Ctx, err error) error {
	res := new(BaseException)
	res.Error = "Internal Server Error"
	res.Message = "Something went wrong"
	res.StatusCode = fiber.StatusInternalServerError

	logger.Error("Global error", err)

	return ctx.Status(int(res.StatusCode)).JSON(res)
}
