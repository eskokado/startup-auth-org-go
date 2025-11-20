package handlers

import (
    "net/http"

    usecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/org"
    "github.com/eskokado/startup-auth-go/backend/pkg/dto"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/gin-gonic/gin"
)

type InviteHandler struct { uc *usecase.InviteMemberUsecase }

func NewInviteHandler(uc *usecase.InviteMemberUsecase) *InviteHandler { return &InviteHandler{uc: uc} }

func (h *InviteHandler) Handle(c *gin.Context) {
    var input dto.InviteParams
    if err := c.ShouldBindJSON(&input); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"}); return }
    orgID, _ := vo.ParseID(input.OrganizationID)
    inviterID, _ := vo.ParseID(input.InviterID)
    inv, err := h.uc.Execute(c, orgID, inviterID, input.Email)
    if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, gin.H{"token": inv.Token})
}