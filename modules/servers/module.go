package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pheethy/go-fiber-tutorial/modules/middleware/handlers"
	"github.com/pheethy/go-fiber-tutorial/modules/middleware/repository"
	"github.com/pheethy/go-fiber-tutorial/modules/middleware/usecase"
	monitorHandler "github.com/pheethy/go-fiber-tutorial/modules/monitor/http"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
}

func InitModuleFactory(router fiber.Router, server *server) IModuleFactory {
	return &moduleFactory{
		router: router,
		server: server,
	}
}

func InitMiddleware(s *server)  handler.ImiddlewareHandler {
	repository := repository.NewMiddlewareRepository(s.db)
	usecase := usecase.NewMiddlewareUsecase(repository)
	return handler.NewMiddlewareHandler(s.cfg, usecase)
}

func (m moduleFactory) MonitorModule() {
	handler := monitorHandler.NewMonitorHandler(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}