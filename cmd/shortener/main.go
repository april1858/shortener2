package main

import (
	chiadapter "github.com/april1858/shortener2/internal/adapter/http/chi"
	"github.com/april1858/shortener2/internal/adapter/repository/memory"
	"github.com/april1858/shortener2/internal/app/shortener"
)

func main() {

	rm := memory.NewRepository()

	service := shortener.NewService(rm)

	useCase := shortener.NewUseCase(service)

	//increment 3
	chiRouter := chiadapter.NewChiHandler(*useCase)
	chiRouter.SetupRoutes()
	chiRouter.Run(":8080")
}
