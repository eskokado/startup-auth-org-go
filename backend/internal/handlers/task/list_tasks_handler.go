package handlers

import (
	"net/http"

	usecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/task"
	"github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
	"github.com/gin-gonic/gin"
)

type ListTasksHandler struct{ uc *usecase.ListTasksUsecase }

func NewListTasksHandler(uc *usecase.ListTasksUsecase) *ListTasksHandler {
	return &ListTasksHandler{uc: uc}
}

func (h *ListTasksHandler) Handle(c *gin.Context) {
	orgID, _ := vo.ParseID(c.Query("organization_id"))
	tasks, err := h.uc.Execute(c, orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
