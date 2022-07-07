package greet

import (
	"Consumer/pkg/logger"
	"net/http"
)

type Handler struct {
	Logger logger.Logger
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.
}
