package handlers

import (
	"HeadHunter/internal/usecases"
)

type Handler struct {
	uc *usecases.UseCases
}

func NewHandler(usecases *usecases.UseCases) *Handler {
	return &Handler{uc: usecases}
}
