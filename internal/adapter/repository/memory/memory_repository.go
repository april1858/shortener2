package memory

import (
	"sync"

	"github.com/april1858/shortener2/internal/app/entity"
)

type Repository struct {
	memory map[string]string
	rwLock *sync.RWMutex
	mx     *sync.RWMutex
}

func NewRepository() *Repository {
	m := make(map[string]string)
	return &Repository{memory: m}
}

func (mr Repository) Store(redirect *entity.Redirect) {
	mr.mx = new(sync.RWMutex)
	mr.mx.Lock()
	mr.memory[redirect.ShortURL] = redirect.OriginalURL
	mr.mx.Unlock()
}

func (mr Repository) Find(code string) string {
	mr.rwLock = new(sync.RWMutex)
	mr.rwLock.Lock()
	defer mr.rwLock.Unlock()
	outURL := mr.memory[code]
	return outURL
}
