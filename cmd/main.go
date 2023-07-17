package main

import (
	"fmt"
	"github.com/Tilvaldiyev/blog-api/internal/app"
	"github.com/Tilvaldiyev/blog-api/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
