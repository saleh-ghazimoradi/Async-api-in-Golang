package store

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"time"
)

type UserStore struct {
	db *sqlx.DB
}

type User struct {
	Id             uuid.UUID `db:"id"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
	CreatedAt      time.Time `db:"created_at"`
}

func (u *User) ComparePassword(password string) error {
	hashedPassword, err := base64.StdEncoding.DecodeString(u.HashedPassword)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func (u *UserStore) CreateUser(ctx context.Context, email, password string) (*User, error) {
	const query = `INSERT INTO users (email, hashed_password) VALUES ($1, $2) RETURNING *;`
	var user User

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(bytes)
	if err := u.db.GetContext(ctx, &user, query, email, hashedPasswordBase64); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (u *UserStore) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	const query = `SELECT * FROM users WHERE email = $1;`
	var user User
	if err := u.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	return &user, nil
}

func (u *UserStore) GetUserById(ctx context.Context, id uuid.UUID) (*User, error) {
	const query = `SELECT * FROM users WHERE id = $1;`
	var user User
	if err := u.db.GetContext(ctx, &user, query, id); err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	return &user, nil
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: sqlx.NewDb(db, "postgres"),
	}
}
