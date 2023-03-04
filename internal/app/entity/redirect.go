package entity

import "errors"

var (
	ErrRedirectNotFound = errors.New("redirect not found")
	ErrRedirectInvalid  = errors.New("invalid url")
)

type Redirect struct {
	OriginalURL string `json:"url" validate:"empty=false & format=url"`
	ShortURL    string `json:"code"`
}
