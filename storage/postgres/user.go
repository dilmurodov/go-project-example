package postgres

import (
	"context"

	"project/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo (db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(ctx context.Context, user *storage.User) (id string, err error) {

	query := `INSERT INTO users (name, age) VALUES($1, $2)`

	err = u.db.QueryRow(ctx, query, user.Name, user.Age).Scan(&id)
	if err != nil{
		return "", err
	}
	
	return id, nil
}

func (u *UserRepo) Get(ctx context.Context, id string) (user *storage.User, err error) {

	query := `SELECT name, age FROM users WHERE id = $1`
	user = &storage.User{}
	err = u.db.QueryRow(ctx, query, id).Scan(
		user.Name,
		user.Age,
	)
	if err != nil{
		return nil, err
	}

	return nil, err
}