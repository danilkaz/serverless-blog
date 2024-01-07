package handlers

import (
	"net/http"
	"serverless-blog/internal/responses"
	"serverless-blog/internal/serializers"
)

type VersionHandler struct {
	version   string
	replicaId string
}

func NewVersionHandler(version, replicaId string) VersionHandler {
	return VersionHandler{version: version, replicaId: replicaId}
}

func (h *VersionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	responses.ResponseData(w, http.StatusOK, serializers.SerializeVersion(h.version, h.replicaId))
}
