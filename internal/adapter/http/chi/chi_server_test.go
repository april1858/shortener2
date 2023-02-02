package chiadapter

import (
	"net/http"
	"testing"

	"github.com/april1858/shortener2/internal/app/shortener"
	"github.com/go-chi/chi/v5"
)

func TestChiHandler_handlerPost(t *testing.T) {
	type fields struct {
		Service   shortener.Service
		ChiRouter *chi.Mux
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ChiHandler{
				Service:   tt.fields.Service,
				ChiRouter: tt.fields.ChiRouter,
			}
			h.handlerPost(tt.args.w, tt.args.r)
		})
	}
}
