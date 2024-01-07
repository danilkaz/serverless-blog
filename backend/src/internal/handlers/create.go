package handlers

import (
	"context"
	"errors"
	"net/http"
	"serverless-blog/internal/deserializers"
	"serverless-blog/internal/models"
	"serverless-blog/internal/responses"
	"serverless-blog/internal/serializers"
)

type PostsCreator interface {
	CreatePost(ctx context.Context, post models.Post) error
}

type CreatePostHandler struct {
	postsCreator PostsCreator
}

func NewCreatePostHandler(postsCreator PostsCreator) CreatePostHandler {
	return CreatePostHandler{postsCreator: postsCreator}
}

func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	post, err := deserializers.DeserializePostRequest(r.Body)
	if err != nil {
		if errors.Is(err, deserializers.ErrUndecodedBody) {
			responses.ResponseError(w, http.StatusBadRequest, err)
		}
		responses.ResponseInternalError(w, err)
	}

	if post.Name == "" {
		responses.ResponseError(w, http.StatusBadRequest, errors.New("post name must be not empty"))
	}

	if post.Text == "" {
		responses.ResponseError(w, http.StatusBadRequest, errors.New("post text must be not empty"))
	}

	if err := h.postsCreator.CreatePost(r.Context(), post); err != nil {
		responses.ResponseInternalError(w, err)
	}

	responses.ResponseData(w, http.StatusCreated, serializers.SerializePost(post))
}
