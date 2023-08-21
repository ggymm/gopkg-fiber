package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/ggymm/gopkg-fiber/log"
)

type Result struct {
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success,omitempty"`
}

func Error(msg string) (r Result) {
	r.Msg = msg
	r.Success = false
	return r
}

func Success(data interface{}) (r Result) {
	r.Data = data
	r.Success = true
	return r
}

type Api struct {
}

func (a *Api) Error(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(Error(msg))
}

func (a *Api) Error400(c *fiber.Ctx, err error) error {
	log.Error().Err(errors.WithStack(err)).Msg("APITraceError400")
	return a.Error(c, http.StatusBadRequest, err.Error())
}

func (a *Api) Error500(c *fiber.Ctx, err error) error {
	log.Error().Err(errors.WithStack(err)).Msg("APITraceError500")
	return a.Error(c, http.StatusInternalServerError, err.Error())
}

func (a *Api) Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Success(data))
}
