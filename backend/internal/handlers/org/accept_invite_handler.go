package handlers

import (
    "net/http"

    usecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/org"
    "github.com/eskokado/startup-auth-go/backend/pkg/dto"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/gin-gonic/gin"
)

type AcceptInviteHandler struct { uc *usecase.AcceptInviteUsecase }

func NewAcceptInviteHandler(uc *usecase.AcceptInviteUsecase) *AcceptInviteHandler { return &AcceptInviteHandler{uc: uc} }

func (h *AcceptInviteHandler) Handle(c *gin.Context) {
    var input dto.AcceptInviteParams
    if err := c.ShouldBindJSON(&input); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"}); return }
    userID, _ := vo.ParseID(input.UserID)
    if err := h.uc.Execute(c, input.Token, userID); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, gin.H{"status": "accepted"})
}