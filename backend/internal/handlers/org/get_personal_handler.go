package handlers

import (
    "net/http"

    "github.com/eskokado/startup-auth-go/backend/pkg/domain/repository"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/gin-gonic/gin"
)

type GetPersonalOrgHandler struct { repo repository.OrganizationRepository }

func NewGetPersonalOrgHandler(repo repository.OrganizationRepository) *GetPersonalOrgHandler { return &GetPersonalOrgHandler{repo: repo} }

func (h *GetPersonalOrgHandler) Handle(c *gin.Context) {
    ownerID, _ := vo.ParseID(c.Param("ownerID"))
    org, err := h.repo.GetByOwnerID(c, ownerID)
    if err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, gin.H{"id": org.ID.String(), "name": org.Name.String()})
}