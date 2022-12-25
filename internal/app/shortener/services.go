package shortener

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/april1858/shortener2/internal/app/entity"
)

type Repository interface {
	Find(code string) string
	Store(redirect *entity.Redirect)
}

type Service struct {
	repository Repository
}

func NewService(redirectRepo Repository) *Service {
	return &Service{
		repository: redirectRepo,
	}
}

func (uc *Service) Find(code string) string {
	return uc.repository.Find(code)
}

func (uc *Service) Store(url string) *entity.Redirect {
	redirect := &entity.Redirect{
		OriginalURL: url,
	}

	redirect.ShortURL = createCode()
	uc.repository.Store(redirect)

	return redirect
}

func createCode() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "1"
	}

	return hex.EncodeToString(b)
}
