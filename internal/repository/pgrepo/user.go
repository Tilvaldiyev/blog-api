package pgrepo

import (
	"context"
	"fmt"
	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/georgysavva/scany/pgxscan"
	"strings"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                hashed_password -- 4
			                )
			VALUES ($1, $2, $3, $4)
			`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUser(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)

	query := fmt.Sprintf("SELECT id, username, first_name, last_name, hashed_password FROM %s WHERE username = $1", usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}

	return user, nil
}
