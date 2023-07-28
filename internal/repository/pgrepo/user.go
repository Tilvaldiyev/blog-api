package pgrepo

import (
	"context"
	"fmt"
	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/georgysavva/scany/pgxscan"
	"strings"
	"time"
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

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUser(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)

	query := fmt.Sprintf("SELECT id, username, first_name, last_name, hashed_password FROM %s WHERE username = $1", usersTable)

	//rows, err := p.SQLDB.Query(query, username)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	err := rows.Scan(&user.ID, &user.Username, &user.LastName, &user.LastName, &user.Password)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//err = rows.Err()
	//if err != nil {
	//	return nil, err
	//}

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}

	return user, nil
}
