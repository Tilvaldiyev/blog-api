package service

import (
	"github.com/Tilvaldiyev/blog-api/internal/config"
	"github.com/Tilvaldiyev/blog-api/internal/repository"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
}

func New(repository repository.Repository, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Config:     config,
	}
}
