package handlers

import (
    "net/http"

    usecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/task"
    "github.com/eskokado/startup-auth-go/backend/pkg/dto"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/gin-gonic/gin"
)

type CreateTaskHandler struct { uc *usecase.CreateTaskUsecase }

func NewCreateTaskHandler(uc *usecase.CreateTaskUsecase) *CreateTaskHandler { return &CreateTaskHandler{uc: uc} }

func (h *CreateTaskHandler) Handle(c *gin.Context) {
    var input dto.CreateTaskParams
    if err := c.ShouldBindJSON(&input); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"}); return }
    orgID, _ := vo.ParseID(input.OrganizationID)
    t, err := h.uc.Execute(c, orgID, input.Title, input.Description)
    if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, gin.H{"id": t.ID.String()})
}