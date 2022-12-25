package memory

import (
	"fmt"

	"github.com/april1858/shortener2/internal/app/entity"
)

type Repository struct {
	memory map[string]string
}

func NewRepository() *Repository {
	m := make(map[string]string)
	return &Repository{memory: m}
}

func (mr Repository) Store(redirect *entity.Redirect) {
	mr.memory[redirect.ShortURL] = redirect.OriginalURL
}

func (mr Repository) Find(code string) string {
	outURL := mr.memory[code]
	fmt.Println("outURL - ", outURL)
	fmt.Println("len memory - ", len(mr.memory))
	return outURL
}