package service

import (
	"github.com/Tilvaldiyev/blog-api/internal/config"
	"github.com/Tilvaldiyev/blog-api/internal/repository"
	"github.com/Tilvaldiyev/blog-api/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Token      *jwttoken.JWTToken
	Config     *config.Config
}

func New(repository repository.Repository, token *jwttoken.JWTToken, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Token:      token,
		Config:     config,
	}
}
