package memory

import (
	"sync"

	"github.com/april1858/shortener2/internal/app/entity"
)

type Repository struct {
	Memory map[string]string
	mx     *sync.RWMutex
}

func NewRepository() *Repository {
	m := make(map[string]string)
	mx := new(sync.RWMutex)
	return &Repository{Memory: m, mx: mx}
}

func (mr Repository) Store(redirect *entity.Redirect) {
	mr.mx.Lock()
	defer mr.mx.Unlock()
	mr.Memory[redirect.ShortURL] = redirect.OriginalURL
}

func (mr Repository) Find(code string) string {
	mr.mx.Lock()
	defer mr.mx.Unlock()
	outURL := mr.Memory[code]
	return outURL
}
