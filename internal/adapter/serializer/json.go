package serializer

import (
	"encoding/json"

	"github.com/april1858/shortener2/internal/app/entity"
)

type JSONSerializer struct{}

func (s *JSONSerializer) Decode(input []byte) (*entity.Redirect, error) {
	redirect := &entity.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, err
	}
	return redirect, nil
}

func (s *JSONSerializer) Encode(input map[string]string) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	return rawMsg, nil
}
