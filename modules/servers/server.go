package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pheethy/go-fiber-tutorial/config"
)

type IServer interface {
	Start()
}

type server struct {
	app *fiber.App
	cfg config.Iconfig
	db  *sqlx.DB
}

func NewServer(cfg config.Iconfig, db *sqlx.DB) IServer {
	return &server{
		cfg: cfg,
		db:  db,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeOut(),
			WriteTimeout: cfg.App().WriteTimeOut(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

func (s *server) Start() {
	// MiddleWare
	mid := InitMiddleware(s)

	// Modules
	v1 := s.app.Group("v1", mid.Cors())
	mod := InitModuleFactory(v1, s)
	mod.MonitorModule()

	// Graceful Shutdown
	var c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func ()  {
		_ = <-c
		log.Println("Server is shutting down...")
		_ = s.app.Shutdown()
	}()

	//Listen to host:port
	log.Printf("Server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
