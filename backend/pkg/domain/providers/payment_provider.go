package providers

import "github.com/eskokado/startup-auth-go/backend/pkg/domain/vo"

type CheckoutRequest struct {
    OrganizationID string
    Plan           vo.PlanType
    Cycle          vo.BillingCycle
    SuccessURL     string
    CancelURL      string
}

type PaymentProvider interface {
    CreateCheckoutSession(req CheckoutRequest) (string, error)
}