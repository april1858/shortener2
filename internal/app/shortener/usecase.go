package shortener

import (
	"github.com/april1858/shortener2/internal/app/entity"
)

type UseCase struct {
	services *Service
}

func NewUseCase(service *Service) *UseCase {
	return &UseCase{
		services: service,
	}
}

func (uc *UseCase) URLToCode(url string) *entity.Redirect {
	return uc.services.Store(url)
}

func (uc *UseCase) CodeToURL(code string) string {
	return uc.services.Find(code)
}
