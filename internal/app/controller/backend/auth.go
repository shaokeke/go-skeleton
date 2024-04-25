package backend

import (
	"log/slog"

	"github.com/MQEnergy/go-skeleton/internal/app/controller"
	"github.com/MQEnergy/go-skeleton/internal/app/service/backend"
	"github.com/MQEnergy/go-skeleton/internal/request/auth"
	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"
)

type AuthController struct {
	controller.Controller
}

var Auth = &AuthController{}

// Index
// @Description: index
// @receiver c
// @param ctx
// @return error
// @author cx
func (c *AuthController) Index(ctx *fiber.Ctx) error {
	params := &auth.IndexReq{}
	if err := c.Validate(ctx, params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "index")
}

// Login 登录
// @Description: login
// @receiver c
// @param ctx
// @return error
// @author cx
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	params := &auth.LoginReq{}
	if !store.Verify(params.CaptchaId, params.Captcha, true) {
		return response.BadRequestException(ctx, "验证码错误")
	}
	if err := c.Validate(ctx, params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := backend.Auth.Login(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}

// Rigster 注册

func (c *AuthController) Rigster(ctx *fiber.Ctx) error {
	params := &auth.RigsterReq{}
	if err := c.Validate(ctx, params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	// service
	info, err := backend.Auth.Rigster(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}

// when set on muti server，use config below，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

func (c *AuthController) Captcha(ctx *fiber.Ctx) error {

	driver := base64Captcha.NewDriverDigit(
		vars.Config.GetInt("captcha.ImgHeight"),
		vars.Config.GetInt("captcha.ImgWidth"),
		vars.Config.GetInt("captcha.KeyLong"),
		0.7,
		80,
	)

	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, _, err := cp.Generate(); err != nil {
		slog.Error("验证码获取失败")
		return response.BadRequestException(ctx, err.Error())
	} else {
		info := fiber.Map{
			"captchaId":     id,
			"picPath":       b64s,
			"captchaLength": vars.Config.GetInt("captcha.KeyLong"),
		}
		return response.SuccessJSON(ctx, "", info)
	}
}
