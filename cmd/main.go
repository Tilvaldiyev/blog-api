package main

import (
	"github.com/Tilvaldiyev/blog-api/internal/app"
	"github.com/Tilvaldiyev/blog-api/internal/config"
)

// @title           Blog API
// @version         0.0.1
// @description     API for Blog application

// @contact.name   Alisher
// @contact.email  email@gmail.com

// @host      localhost:8080
// @BasePath  /api/v2

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
