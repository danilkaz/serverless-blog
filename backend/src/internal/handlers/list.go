package handlers

import (
	"context"
	"net/http"
	"serverless-blog/internal/models"
	"serverless-blog/internal/responses"
	"serverless-blog/internal/serializers"
)

type PostsProvider interface {
	Posts(ctx context.Context) ([]models.Post, error)
}

type ListPostsHandler struct {
	postsProvider PostsProvider
}

func NewListPostsHandler(postsProvider PostsProvider) ListPostsHandler {
	return ListPostsHandler{postsProvider: postsProvider}
}

func (h *ListPostsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postsProvider.Posts(r.Context())
	if err != nil {
		responses.ResponseInternalError(w, err)
		return
	}

	responses.ResponseData(w, http.StatusOK, serializers.SerializeListPosts(posts))
}
