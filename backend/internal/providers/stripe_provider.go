package providers

import (
    "os"
    "errors"
    "github.com/eskokado/startup-auth-go/backend/pkg/domain/providers"
)

type StripeProvider struct{}

func NewStripeProvider() *StripeProvider { return &StripeProvider{} }

func (s *StripeProvider) CreateCheckoutSession(req providers.CheckoutRequest) (string, error) {
    if os.Getenv("STRIPE_SECRET") == "" { return "", errors.New("stripe not configured") }
    // Placeholder: integração real com stripe-go deve ser adicionada
    return os.Getenv("APP_BASE_URL") + "/billing/success", nil
}