package main

import (
	chiadapter "github.com/april1858/shortener2/internal/adapter/http/chi"
	"github.com/april1858/shortener2/internal/adapter/repository/memory"
	"github.com/april1858/shortener2/internal/app/shortener"
	"github.com/april1858/shortener2/config"
)

func main() {
	cnfg := config.NewConfig()

	mr := memory.NewRepository()

	service := shortener.NewService(mr)

	chiRouter := chiadapter.NewChiHandler(*service)
	chiRouter.SetupRoutes()
	chiRouter.Run(":" + cnfg.ServerAdres)
}
