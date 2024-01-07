package serializers

import (
	"serverless-blog/internal/models"
	"time"
)

type VersionResponse struct {
	Version   string `json:"version"`
	ReplicaId string `json:"replica_id"`
}

type PostResponse struct {
	Name      string    `json:"name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type ListPostsResponse struct {
	Items []PostResponse `json:"items"`
}

func SerializeVersion(version, replicaId string) VersionResponse {
	return VersionResponse{
		Version:   version,
		ReplicaId: replicaId,
	}
}

func SerializePost(post models.Post) PostResponse {
	return PostResponse{
		Name:      post.Name,
		Text:      post.Text,
		CreatedAt: post.CreatedAt,
	}
}

func SerializeListPosts(posts []models.Post) ListPostsResponse {
	items := make([]PostResponse, 0, len(posts))

	for _, post := range posts {
		items = append(items, SerializePost(post))
	}

	return ListPostsResponse{Items: items}
}
