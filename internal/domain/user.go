package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]User, string, error)
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, user *User) error
	GetByTitle(ctx context.Context, title string) (User, error)
	Store(context.Context, *User) error
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []User, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (User, error)
	GetByTitle(ctx context.Context, title string) (User, error)
	Update(ctx context.Context, user *User) error
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}