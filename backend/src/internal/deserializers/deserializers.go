package deserializers

import (
	"encoding/json"
	"errors"
	"io"
	"serverless-blog/internal/models"
	"time"
)

var (
	ErrUndecodedBody error = errors.New("can't decode request body")
)

type CreatePostRequest struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func DeserializePostRequest(body io.Reader) (models.Post, error) {
	var request CreatePostRequest

	if err := json.NewDecoder(body).Decode(&request); err != nil {
		return models.Post{}, ErrUndecodedBody
	}

	return models.Post{
		Name:      request.Name,
		Text:      request.Text,
		CreatedAt: time.Now(),
	}, nil
}
