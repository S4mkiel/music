package http

import (
	_ "github.com/S4mkiel/music/docs"
	controller "github.com/S4mkiel/music/infra/http/controllers"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("http",
	FiberModule,
	fx.Invoke(RegisterControllers),
	fx.Provide(controller.NewMusicController),
)

func RegisterControllers(app *fiber.App, musicController *controller.MusicController) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/swagger/*", swagger.HandlerDefault) // default
	v1.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "http://example.com/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	musicController.RegisterRoutes(v1)
}
