package chiadapter

import (
	"encoding/json"
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
	Service   shortener.Service
	ChiRouter *chi.Mux
}

func NewChiHandler(services shortener.Service) *ChiHandler {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)
	return &ChiHandler{
		Service:   services,
		ChiRouter: chiRouter,
	}
}

func (h *ChiHandler) Run(address string) {
	fmt.Printf("CHI  listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, h.ChiRouter))
}

func (h *ChiHandler) SetupRoutes() {
	h.ChiRouter.Get("/{code}", h.handlerGet)
	h.ChiRouter.Post("/", h.handlerPost)
	h.ChiRouter.Post("/api/shorten", h.handlerJSON)
}

func (h *ChiHandler) handlerGet(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	redirect := h.Service.Find(code)
	status := http.StatusTemporaryRedirect
	if redirect == "" {
		status = http.StatusBadRequest
	}

	http.Redirect(w, r, redirect, status)
}

func (h *ChiHandler) handlerPost(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "http.StatusBadRequest", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	redirect := h.Service.Store(string(requestBody))

	answer := "http://" + host + "/" + redirect.ShortURL
	returnResponse(w, contentType, []byte(answer), status)
}

func (h *ChiHandler) handlerJSON(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	if host == "" {
		host = "localhost:8080"
	}
	contentType := "application/json"

	status := http.StatusCreated
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	redirect := &entity.Redirect{}
	if err := json.Unmarshal(requestBody, redirect); err != nil {
		return
	}

	_, err = url.ParseRequestURI(redirect.OriginalURL)
	if err != nil {
		http.Error(w, "http.statusBadRequest", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	redirect = h.Service.Store(redirect.OriginalURL)

	answerStruct := map[string]string{"result":"http://" + host + "/" + redirect.ShortURL}
	answer, err := json.Marshal(answerStruct)
	if err != nil {
		return
	}
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
