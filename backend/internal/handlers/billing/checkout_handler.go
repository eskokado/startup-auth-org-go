package handlers

import (
    "net/http"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/providers"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"
    "github.com/eskokado/startup-auth-go/backend/pkg/dto"
    "github.com/gin-gonic/gin"
)

type CheckoutHandler struct { payment providers.PaymentProvider }

func NewCheckoutHandler(p providers.PaymentProvider) *CheckoutHandler { return &CheckoutHandler{payment: p} }

func (h *CheckoutHandler) Handle(c *gin.Context) {
    var input dto.CheckoutParams
    if err := c.ShouldBindJSON(&input); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"}); return }
    plan, err := vo.NewPlanType(input.Plan)
    if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    cycle, err := vo.NewBillingCycle(input.Cycle)
    if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    url, err := h.payment.CreateCheckoutSession(providers.CheckoutRequest{OrganizationID: input.OrganizationID, Plan: plan, Cycle: cycle, SuccessURL: input.SuccessURL, CancelURL: input.CancelURL})
    if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, gin.H{"checkout_url": url})
}