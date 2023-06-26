package servers

import (
	"encoding/json"

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

func (s server) Start() {
	// Graceful Shutdown
}
