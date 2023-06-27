package main

import (
	"context"
	"os"

	"github.com/pheethy/go-fiber-tutorial/config"
	"github.com/pheethy/go-fiber-tutorial/modules/servers"
	"github.com/pheethy/go-fiber-tutorial/pkg/database"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	var ctx = context.Background()
	var cfg = config.LoadConfig(envPath())
	var db = database.DbConnect(ctx, cfg.Db())
	defer db.Close()

	var serve = servers.NewServer(cfg, db)
	serve.Start()
}
