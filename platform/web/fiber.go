package web

import (
	"fmt"
	"time"

	config "github.com/KornCode/KUKR-APIs-Service/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var fb fiberInstance

type fiberInstance struct {
	App *fiber.App
}

func NewFiberApp(conf config.Server) *fiber.App {
	const idleTimeout = 5 * time.Second

	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Use(recover.New())

	fb = fiberInstance{
		App: app,
	}

	return fb.App
}

func ListenFiberApp(conf config.Server) error {
	if err := fb.App.Listen(fmt.Sprintf(":%s", conf.Port)); err != nil {
		return err
	}

	return nil
}

func ShutdownFiberApp() error {
	if err := fb.App.Shutdown(); err != nil {
		return err
	}

	return nil
}
