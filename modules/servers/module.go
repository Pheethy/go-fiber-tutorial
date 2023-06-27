package servers

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/pheethy/go-fiber-tutorial/modules/monitor/http"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
}

func NewModuleFactory(router fiber.Router, server *server) IModuleFactory {
	return &moduleFactory{
		router: router,
		server: server,
	}
}

func (m moduleFactory) MonitorModule() {
	handler := handler.NewMonitorHandler(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}