package handlers

import (
    "net/http"

    usecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/task"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/entity"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/eskokado/startup-auth-go/backend/pkg/dto"
    "github.com/gin-gonic/gin"
)

type UpdateTaskStatusHandler struct { uc *usecase.UpdateTaskStatusUsecase }

func NewUpdateTaskStatusHandler(uc *usecase.UpdateTaskStatusUsecase) *UpdateTaskStatusHandler { return &UpdateTaskStatusHandler{uc: uc} }

func (h *UpdateTaskStatusHandler) Handle(c *gin.Context) {
    var input dto.UpdateTaskStatusParams
    if err := c.ShouldBindJSON(&input); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"}); return }
    // Para simplificar, reconstruímos Task com dados mínimos
    id, _ := vo.ParseID(input.TaskID)
    orgID, _ := vo.ParseID(input.OrganizationID)
    t := &entity.Task{ID: id, OrganizationID: orgID, Title: input.Title, Description: input.Description}
    updated, err := h.uc.Execute(c, t, input.Status)
    if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, gin.H{"status": updated.Status.String()})
}