package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pheethy/go-fiber-tutorial/config"
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
	cfg := config.LoadConfig(envPath())
	db := database.DbConnect(ctx, cfg.Db())
	defer db.Close()
	fmt.Println(os.Args)
}
