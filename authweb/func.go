package authweb

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	auth "github.com/ggymm/gopkg-auth"
	"github.com/ggymm/gopkg/utils"
)

func Login(id int64, ctx *fiber.Ctx, config ...auth.LoginConfig) (string, error) {
	if auth.NotInit() {
		return "", errors.New(auth.ErrAuthNotInit)
	}

	// 获取配置
	var cfg = auth.LoginConfig{
		Device:  "web",
		Timeout: auth.GetDefaultTimeout(),
	}
	if len(config) > 0 {
		cfg = config[0]
	}

	// 执行登陆
	token, err := auth.Login(id, cfg)
	if err != nil {
		return "", err
	}

	// 写入 cookie
	cookie := fasthttp.AcquireCookie()
	cookie.SetKey(auth.GetTokenName())
	cookie.SetValue(token)
	cookie.SetPath("/")
	// cookie.SetDomain(ctx.Hostname()) // 会自动设置
	if cfg.Timeout == -1 {
		// 设置永不过期
		cookie.SetMaxAge(utils.YearToSecond(1))
		// cookie.SetExpire(time.Now().AddDate(1, 0, 0))
	} else if cfg.Timeout > 0 {
		cookie.SetMaxAge(int(cfg.Timeout.Seconds()))
		// cookie.SetExpire(time.Now().Add(time.Duration(cfg.Timeout) * time.Second))
	}
	ctx.Response().Header.SetCookie(cookie)
	fasthttp.ReleaseCookie(cookie)

	// 写入 response header
	ctx.Response().Header.Set(auth.GetTokenName(), token)
	return token, nil
}

func Check(ctx *fiber.Ctx) (bool, error) {
	if auth.NotInit() {
		return false, errors.New(auth.ErrAuthNotInit)
	}
	tokenName := auth.GetTokenName()

	// 从请求体中获取 token
	token := ctx.Get(tokenName)

	// 从请求头中获取 token
	if len(token) == 0 {
		token = ctx.GetReqHeaders()[tokenName]
	}

	// 从 cookie 中获取 token
	if len(token) == 0 {
		token = ctx.Cookies(tokenName)
	}

	return auth.Check(token)
}

func GetSession(ctx *fiber.Ctx) (interface{}, error) {
	if auth.NotInit() {
		return nil, errors.New(auth.ErrAuthNotInit)
	}
	tokenName := auth.GetTokenName()

	// 从请求体中获取 token
	token := ctx.Get(tokenName)

	// 从请求头中获取 token
	if len(token) == 0 {
		token = ctx.GetReqHeaders()[tokenName]
	}

	// 从 cookie 中获取 token
	if len(token) == 0 {
		token = ctx.Cookies(tokenName)
	}

	// 获取 session
	data, err := auth.GetSession(token)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SaveSession(id int64, value interface{}) error {
	if auth.NotInit() {
		return errors.New(auth.ErrAuthNotInit)
	}
	return auth.SaveSession(id, value)
}
