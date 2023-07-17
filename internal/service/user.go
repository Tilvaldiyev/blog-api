package service

import (
	"context"
	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/Tilvaldiyev/blog-api/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
