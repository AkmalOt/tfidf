package handler

import (
	"github.com/gin-gonic/gin"
	"tdidf/internal/service"
)

type Handler struct {
	Engine  *gin.Engine
	Service *service.ServiceImpl
}

func NewHandler(engine *gin.Engine, service *service.ServiceImpl) *Handler {
	return &Handler{
		Engine:  engine,
		Service: service,
	}
}

func (h *Handler) Init() {
	h.Engine.POST("/tfidf", h.TfIdf)
}
