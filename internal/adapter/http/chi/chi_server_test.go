package chiadapter

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/april1858/shortener2/internal/adapter/repository/memory"
	"github.com/april1858/shortener2/internal/app/shortener"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Test_handlers(t *testing.T) {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)
	tt := []struct {
		name        string
		originalURL string
		shortURL    string
		status      int
	}{
		{name: "ok", originalURL: "http://s-p.ru", shortURL: "", status: http.StatusCreated},
		{name: "!ok", originalURL: "saryg.ru", shortURL: "", status: http.StatusBadRequest},
	}

	rm := memory.NewRepository()
	service := shortener.NewService(rm)
	useCase := shortener.NewUseCase(service)

	for _, tc := range tt {
		req, _ := http.NewRequest(http.MethodPost, "localhost:8080/", strings.NewReader(tc.originalURL))
		rec := httptest.NewRecorder()
		handlerPost(rec, req, useCase)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.status {
			t.Errorf("expected status %v got %v", res.StatusCode, tc.status)
		}

		answer, _ := io.ReadAll(res.Body)
		tc.shortURL = string(answer)
		if tc.shortURL != "http.StatusBadRequest" {
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("code", tc.shortURL[len(tc.shortURL)-8:])
			if tc.status != http.StatusBadRequest {
				r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/"+tc.shortURL, nil)
				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				rec := httptest.NewRecorder()
				handlerGet(rec, r, useCase)
				res := rec.Result()
				defer res.Body.Close()
				if res.StatusCode != http.StatusTemporaryRedirect {
					t.Errorf("expected status %v got %v", res.StatusCode, http.StatusTemporaryRedirect)
				}
			}

		}
	}
}
