package handlers

import (
    "net/http"

    "github.com/eskokado/startup-auth-go/backend/internal/usecase/task"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/gin-gonic/gin"
)

type DeleteTaskHandler struct { uc *usecase.DeleteTaskUsecase }

func NewDeleteTaskHandler(uc *usecase.DeleteTaskUsecase) *DeleteTaskHandler { return &DeleteTaskHandler{uc: uc} }

func (h *DeleteTaskHandler) Handle(c *gin.Context) {
    id, _ := vo.ParseID(c.Param("id"))
    if err := h.uc.Execute(c, id); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, gin.H{"deleted": true})
}