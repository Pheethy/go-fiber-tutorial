package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pheethy/go-fiber-tutorial/config"
	"github.com/pheethy/go-fiber-tutorial/models"
	"github.com/pheethy/go-fiber-tutorial/modules/monitor"
)

type monitorHandler struct {
	cfg config.Iconfig
}

func NewMonitorHandler(cfg config.Iconfig) monitor.IMonitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	resp := models.Monitor{
		Name: h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return c.Status(http.StatusOK).JSON(resp)
}