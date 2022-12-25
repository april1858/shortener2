package chiadapter

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/april1858/shortener2/internal/app/entity"
	"github.com/april1858/shortener2/internal/app/shortener"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ChiHandler struct {
	UseCase   shortener.UseCase
	ChiRouter *chi.Mux
}

func NewChiHandler(useCase shortener.UseCase) *ChiHandler {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)
	return &ChiHandler{
		UseCase:   useCase,
		ChiRouter: chiRouter,
	}
}

func (h *ChiHandler) Run(address string) {
	fmt.Printf("CHI  listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, h.ChiRouter))
}

func (h *ChiHandler) SetupRoutes() {
	h.ChiRouter.Get("/{code}", func(w http.ResponseWriter, r *http.Request) {
		handlerGet(w, r, &h.UseCase)
	})
	h.ChiRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlerPost(w, r, &h.UseCase)
	})
}

func handlerGet(w http.ResponseWriter, r *http.Request, useCase *shortener.UseCase) {

	code := chi.URLParam(r, "code")
	redirect := useCase.CodeToURL(code)
	status := http.StatusTemporaryRedirect
	if redirect == "" {
		status = http.StatusBadRequest
	}

	http.Redirect(w, r, redirect, status)
}

func handlerPost(w http.ResponseWriter, r *http.Request, useCase *shortener.UseCase) {

	host := r.Host
	if host == "" {
		host = "localhost:8080"
	}
	contentType := r.Header.Get("Content-Type")
	status := http.StatusCreated
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, err = url.ParseRequestURI(string(requestBody))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	redirectOut := useCase.URLToCode(string(requestBody))
	if err != nil {
		if err == entity.ErrRedirectInvalid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	answer := "http://" + host + "/" + redirectOut.ShortURL + "\n"
	returnResponse(w, contentType, []byte(answer), status)
}

func returnResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
