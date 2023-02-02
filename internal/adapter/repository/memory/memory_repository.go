package memory

import (
	"sync"

	"github.com/april1858/shortener2/internal/app/entity"
)

type Repository struct {
	memory map[string]string
	mx     *sync.RWMutex
}

func NewRepository() *Repository {
	m := make(map[string]string)
	return &Repository{memory: m}
}

func (mr Repository) Store(redirect *entity.Redirect) {
	mr.mx = new(sync.RWMutex)
	mr.mx.Lock()
	defer mr.mx.Unlock()
	mr.memory[redirect.ShortURL] = redirect.OriginalURL
}

func (mr Repository) Find(code string) string {
	mr.mx = new(sync.RWMutex)
	mr.mx.Lock()
	defer mr.mx.Unlock()
	outURL := mr.memory[code]
	return outURL
}
