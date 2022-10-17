package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/usecases"
)

type Handler struct {
	uc  *usecases.UseCases
	Cfg configs.Config
}

func NewHandler(usecases *usecases.UseCases) *Handler {
	return &Handler{uc: usecases, Cfg: usecases.Cfg}
}
