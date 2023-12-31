package logger

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/ggymm/gopkg-fiber/log"
)

func needLog(contentType []byte) bool {
	var typeList = []string{
		"application/json",
		"application/xml",
		"text/xml",
		"text/plain",
		"application/x-www-form-urlencoded",
	}

	// 判断请求类型 是否在 contentTypes 中
	for _, t := range typeList {
		if string(contentType) == t {
			return true
		}
	}
	return false
}

func NewLogger() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		// next middleware
		if err := ctx.Next(); err != nil {
			return err
		}

		// 计算请求耗时
		elapsed := time.Since(start)
		costTime := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)

		var params map[string]string
		var reqBody []byte

		// 判断请求类型 是否在 contentTypes 中

		if needLog(ctx.Request().Header.ContentType()) {
			params = ctx.Queries()
			reqBody = ctx.Request().Body()
		}

		var respBody []byte
		if needLog(ctx.Response().Header.ContentType()) {
			respBody = ctx.Response().Body()
		}

		// 记录请求日志
		log.Info().
			// 请求基本参数
			Str("ip", ctx.IP()).
			Str("path", ctx.Path()).
			Str("method", ctx.Method()).
			Str("costTime", costTime).

			// 请求参数
			Str("params", fmt.Sprintf("%v", params)).
			Str("reqBody", fmt.Sprintf("%s", reqBody)).

			// 响应参数
			Str("respBody", fmt.Sprintf("%s", respBody)).
			Msg("APITrace")

		return nil
	}
}
