package database

import (
	"context"
	"log"
	
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pheethy/go-fiber-tutorial/config"
)

func DbConnect(ctx context.Context,cfg config.IDbConfig) *sqlx.DB {
	db, err := sqlx.ConnectContext(ctx, "pgx", cfg.Url())
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxConns())

	return db
}