package backend

import (
    "github.com/MQEnergy/go-skeleton/internal/app/controller"
    "github.com/MQEnergy/go-skeleton/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ArticleController struct{
	controller.Controller
}

var Article = &ArticleController{}

// Index ...
func (s *ArticleController) Index(ctx *fiber.Ctx) error {
    // Todo implement ...
	return response.SuccessJSON(ctx, "", "")
}