package main

import (
	chiadapter "github.com/april1858/shortener2/internal/adapter/http/chi"
	"github.com/april1858/shortener2/internal/adapter/repository/memory"
	"github.com/april1858/shortener2/internal/app/shortener"
)

func main() {

	mr := memory.NewRepository()

	service := shortener.NewService(mr)

	chiRouter := chiadapter.NewChiHandler(*service)
	chiRouter.SetupRoutes()
	chiRouter.Run(":8080")
}
