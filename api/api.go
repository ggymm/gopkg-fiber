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
	Success bool        `json:"success"`
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

func (a *Api) Error(ctx *fiber.Ctx, status int, msg string) error {
	return ctx.Status(status).JSON(Error(msg))
}

func (a *Api) Error400(ctx *fiber.Ctx, err error) error {
	log.Error().Err(errors.WithStack(err)).Msg("APITraceError400")
	return a.Error(ctx, http.StatusBadRequest, err.Error())
}

func (a *Api) Error500(ctx *fiber.Ctx, err error) error {
	log.Error().Err(errors.WithStack(err)).Msg("APITraceError500")
	return a.Error(ctx, http.StatusInternalServerError, err.Error())
}

func (a *Api) Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(Success(data))
}
