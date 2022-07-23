package storage

import "context"

type StorageI interface {
	CloseDB()
	User() UserI
}

type User struct {
	Id   string
	Name string
	Age  int64
}

type UserI interface {
	Get(ctx context.Context, id string) (user *User, err error)
	Create(ctx context.Context, user *User) (id string, err error)
}
