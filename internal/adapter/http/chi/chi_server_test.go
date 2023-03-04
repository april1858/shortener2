package chiadapter

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/april1858/shortener2/internal/adapter/repository/memory"
	"github.com/april1858/shortener2/internal/app/shortener"
)

func TestChiHandler_handlerPost(t *testing.T) {
	mr := memory.NewRepository()

	service := shortener.NewService(mr)

	ChiHandler := NewChiHandler(*service)

	tests := []struct {
		name        string
		status      int
		originalURL string
	}{
		{name: "first", status: http.StatusCreated, originalURL: "http://yandex-practicum.ru"},
		{name: "second", status: http.StatusBadRequest, originalURL: "p//yandex-practicum.ru"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", strings.NewReader(tt.originalURL))
			ChiHandler.handlerPost(wr, req)
			res := wr.Result()
			res.Body.Close()
			statusCode := res.StatusCode
			if statusCode != tt.status {
				t.Errorf("StatusCode = %v, want %v", statusCode, tt.status)
			}
		})
	}
}

/*
func TestChiHandler_handlerGet(t *testing.T) {
	mr := memory.NewRepository()
	mr.Memory["12345678"] = "http://yandex-practicum.ru"
	service := shortener.NewService(mr)

	ChiHandler := NewChiHandler(*service)

	tests := []struct {
		name     string
		status   int
		shortURL string
	}{
		{name: "first", status: http.StatusTemporaryRedirect, shortURL: "12345678"},
		{name: "second", status: http.StatusBadRequest, shortURL: "22345678"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/"+tt.shortURL, nil)

			ChiHandler.handlerGet(wr, req)        // chi.URLParam(r, "code") - ?
			if wr.Result().StatusCode != tt.status {
				t.Errorf("StatusCode = %v, want %v", wr.Result().StatusCode, tt.status)
			}
		})
	}
}
*/
