package postgres

import (
	"context"
	"database/sql"
	"strconv"

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

	query := `INSERT INTO user_data (name, age) VALUES($1, $2) returning user_id`

	err = u.db.QueryRow(ctx, query, user.Name, user.Age).Scan(&id)
	if err != nil{
		return "", err
	}
	
	return id, nil
}

func (u *UserRepo) Get(ctx context.Context, id string) (user *storage.User, err error) {

	var int_id sql.NullInt32

	query := `SELECT id, name, age FROM user_data WHERE user_id = $1`
	user = &storage.User{}
	err = u.db.QueryRow(ctx, query, id).Scan(
		&int_id,
		&user.Name,
		&user.Age,
	)
	if err != nil{
		return nil, err
	}

	if int_id.Valid{
		user.Id = strconv.Itoa(int(int_id.Int32))
	}

	return user, nil
}