package postgres

import (
	"context"
	"fmt"
	"project/storage"

	"project/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db   *pgxpool.Pool
	user storage.UserI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage storage.StorageI, err error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
	cfg.PostgresUser,
	cfg.PostgresPassword,
	cfg.PostgresHost,
	cfg.PostgresPort,
	cfg.PostgresDatabase))
	if err != nil{
		return nil, err
	}

	config.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil{
		return nil, err
	}

	return  &Store{
		db: pool,
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserI {
	if s.user == nil{
		s.user = NewUserRepo(s.db)
	}
	
	return s.user
}