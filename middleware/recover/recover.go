package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/ggymm/gopkg-fiber/log"
)

func NewRecover() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {

		defer func() {
			if r := recover(); r != nil {

				log.Error().Err(errors.WithStack(r.(error))).Msg("服务器内部错误")

				err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"msg":     "服务器内部错误",
					"success": false,
				})
			}
		}()

		return c.Next()
	}
}
